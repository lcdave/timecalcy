<template>
  <div>
    <div class="flex justify-between w-96 mb-6">
      <div>
        <label for="startTime" class="block mb-[10px]">Start</label>
        <input
          v-model="startTime"
          name="start"
          type="time"
          class="py-1 px-4 rounded-xl block"
        />
      </div>
      <div>
        <label for="endTime" class="block mb-[10px]">End</label>
        <input
          v-model="endTime"
          name="end"
          type="time"
          class="py-1 px-4 rounded-xl block"
        />
      </div>
      <div>
        <label for="breakTime" class="block mb-[10px]">Break</label>
        <input
          v-model="breakTime"
          name="breakTime"
          type="time"
          class="py-1 px-4 rounded-xl block"
        />
      </div>
    </div>
    <div>
      <label for="result" class="block mb-[10px]">Worked</label>
      <div>{{ result }}</div>
    </div>
    <button @click="saveEntry" class="px-4 py-2 rounded-xl bg-blue-500 text-white">
      Save
    </button>
  </div>
</template>

<script>
export default {
  data() {
    return {
      startTime: "",
      endTime: "",
      breakTime: "",
      result: "",
    };
  },
  watch: {
    startTime: function () {
      this.calculate();
    },
    endTime: function () {
      this.calculate();
    },
    breakTime: function () {
      this.calculate();
    },
  },
  methods: {
    calculate() {
      let start = this.startTime.split(":");
      let end = this.endTime.split(":");
      let breakTm = this.breakTime.split(":");

      const startSeconds =
        this.hoursToSeconds(start[0]) + this.minutesToSeconds(start[1]);

      const breakSeconds =
        this.hoursToSeconds(breakTm[0]) + this.minutesToSeconds(breakTm[1]);
      const endSeconds =
        this.hoursToSeconds(end[0]) + this.minutesToSeconds(end[1]);

      let difference = endSeconds - startSeconds - breakSeconds;

      this.result = this.secondsToHoursAndMinutes(difference);
    },
    hoursToSeconds(hours) {
      return hours * 3600;
    },
    minutesToSeconds(minutes) {
      return minutes * 60;
    },
    secondsToHoursAndMinutes(seconds) {
      let hours = Math.floor(seconds / 3600);
      let minutes = Math.floor((seconds - hours * 3600) / 60);

      return `${hours}h ${minutes}m`;
    },
    async saveEntry() {
      await $fetch( 'https://localhost:8080/workentry/create', {
        method: 'POST',
        body: "huhu"
      } );
    }
  },
};
</script>

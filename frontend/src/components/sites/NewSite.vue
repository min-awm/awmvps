<template>
  <div class="m-2 bg-white border rounded-md p-9">
    <div class="pb-5">
      <button
        v-for="(tab, index) in tabs"
        :key="`NewSite-tabs-${index}`"
        @click="activeTab = index"
        class="flex-1 px-4 py-3 text-sm font-medium text-center transition-colors duration-200"
        :class="
          activeTab === index
            ? 'text-blue-600 border-b-2 border-blue-500'
            : 'text-gray-500 hover:text-gray-700 '
        "
        role="tab"
      >
        {{ tab }}
      </button>
    </div>

    <div v-if="activeTab === 0" class="flex flex-col gap-3">
      <component
        v-for="(comp, i) in serviceComps"
        :is="comp"
        :key="`NewSite-serviceComps-${i}`"
      />

      <SelecteService @selected="selectedServiceHandle" />
      <!-- <button
        class="items-center inline px-4 py-2 text-sm font-medium text-white bg-blue-600 border border-transparent rounded-md shadow-sm max-w-32 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
      >
        Thêm service
      </button> -->
    </div>
    <div v-if="activeTab === 1">
      <textarea
        id="config-content"
        rows="20"
        class="block w-full p-2 mt-1 font-mono border border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
        v-model="dockerComposeText"
      ></textarea>
    </div>
  </div>
</template>

<script setup>
import { ref, shallowRef } from "vue";
import yaml from "js-yaml";
import NodejsComp from "@/components/sites/services/Nodejs.vue";
import SelecteService from "@/components/sites/SelecteService.vue";
const serviceComps = shallowRef([NodejsComp]);
const tabs = [
  "Cấu hình",
  // "docker-compose.yml"
];

const activeTab = ref(0);
const dockerComposeText = ref("");

function yamlToJson() {
  try {
    const data = yaml.load(dockerComposeText.value);
    console.log(data); // { name: 'John Doe', age: 30, skills: ['JavaScript', 'Python', 'YAML'] }
  } catch (e) {
    console.error("Error parsing YAML:", e);
  }
}

function selectedServiceHandle(val) {
  switch (val.name) {
    case "Nodejs":
      serviceComps.value = [...serviceComps.value, NodejsComp];
      break;

    default:
      break;
  }
}
</script>

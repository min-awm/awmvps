<template>
  <div class="flex flex-col flex-1 overflow-hidden">
    <div class="p-4 sm:p-6">
      <div class="flex items-center gap-2 mb-6">
        <button
          class="flex items-center text-sm text-blue-600 hover:text-blue-700 md:hidden"
          @click="$emit('openSideBarFn')"
        >
          <ChevronRight class="mr-1" />
        </button>
        <h1 class="text-xl font-semibold text-gray-900">Containers</h1>
      </div>

      <!-- Container Table -->
      <div class="px-2 overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th
                scope="col"
                class="px-3 py-3 text-xs font-medium tracking-wider text-left text-gray-500 uppercase"
              >
                Name
              </th>
              <th
                scope="col"
                class="px-3 py-3 text-xs font-medium tracking-wider text-left text-gray-500 uppercase"
              >
                Container ID
              </th>
              <th
                scope="col"
                class="px-3 py-3 text-xs font-medium tracking-wider text-left text-gray-500 uppercase"
              >
                Image
              </th>
              <th
                scope="col"
                class="px-3 py-3 text-xs font-medium tracking-wider text-left text-gray-500 uppercase"
              >
                Network
              </th>
              <th
                scope="col"
                class="px-3 py-3 text-xs font-medium tracking-wider text-left text-gray-500 uppercase"
              >
                Port(s)
              </th>
              <th
                scope="col"
                class="px-3 py-3 text-xs font-medium tracking-wider text-left text-gray-500 uppercase"
              >
                Status
              </th>
              <th
                scope="col"
                class="px-3 py-3 text-xs font-medium tracking-wider text-right text-gray-500 uppercase"
              >
                Actions
              </th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="(container, i) in containers" :key="`Container-${i}`">
              <td class="px-3 py-4">
                <div class="flex items-center">
                  {{ container.Names }}
                </div>
              </td>
              <td class="px-3 py-4 text-sm text-gray-500">
                {{ container.ID }}
              </td>
              <td class="px-3 py-4 text-sm text-gray-500">
                {{ container.Image }}
              </td>

              <td class="px-3 py-4 text-sm text-gray-500">
                {{ container.Networks }}
              </td>
              <td class="px-3 py-4 text-sm text-gray-500">
                {{ container.Ports }}
              </td>
              <td class="px-3 py-4 text-sm text-gray-500">
                {{ container.Status }}
              </td>
              <td
                class="flex items-center justify-end gap-3 px-3 py-4 space-x-2 text-right"
              >
                <button
                  class="text-gray-400 hover:text-gray-500"
                  v-if="container.Status.substring(0, 2) == 'Up'"
                  @click="stopContainer(container.ID)"
                >
                  <OctagonPause class="w-4 h-4" />
                </button>
                <button
                  class="text-gray-400 hover:text-gray-500"
                  v-else
                  @click="startContainer(container.ID)"
                >
                  <Play class="w-4 h-4" />
                </button>

                <button
                  class="text-gray-400 hover:text-gray-500"
                  @click="removeContainer(container.ID)"
                >
                  <Trash2 class="w-4 h-4" />
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <div class="p-6">
      <h2 class="mb-4 text-2xl font-semibold">Logs</h2>
      <div
        class="h-48 p-4 overflow-y-auto font-mono text-sm text-green-400 bg-gray-900 rounded-lg"
      >
        <p v-for="(log, i) in dataLogs" :key="`Container-log-${i}`">
          {{ log }}
        </p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { Play, Trash2, ChevronRight, OctagonPause } from "lucide-vue-next";
import axios from "@/axios";
import API from "@/api";

const dataLogs = ref([]);
const containers = ref([]);
getListContainer();

async function getListContainer() {
  try {
    const res = await axios.get(`${API.LIST_CONTAINER_DOCKER}`);
    if (res.success) {
      if (res.message == "") {
        containers.value = [];
      } else {
        const jsonStrings = res.message
          .trim()
          .split("\n")
          .map((line) => line.trim())
          .map((line) => line.slice(1, -1));

        const parsedData = jsonStrings.map(JSON.parse);
        containers.value = parsedData;
      }
    } else {
      dataLogs.value.push(
        `[${new Date().toLocaleTimeString()}] ${res.message}`
      );
    }
  } catch (error) {
    console.log(`Docker: ${error}`);
  }
}

async function stopContainer(id) {
  dataLogs.value.push(
    `[${new Date().toLocaleTimeString()}] Đang dừng container ${id}`
  );
  try {
    const formData = new FormData();
    formData.append("id", id);
    const res = await axios.post(`${API.STOP_CONTAINER_DOCKER}`, formData);
    if (res.success) {
      getListContainer();
    }
    dataLogs.value.push(`[${new Date().toLocaleTimeString()}] ${res.message}`);
  } catch (error) {
    console.log(`Docker: ${error}`);
  }
}

async function startContainer(id) {
  dataLogs.value.push(
    `[${new Date().toLocaleTimeString()}] Đang chạy container ${id}`
  );
  try {
    const formData = new FormData();
    formData.append("id", id);
    const res = await axios.post(`${API.START_CONTAINER_DOCKER}`, formData);
    if (res.success) {
      getListContainer();
    }
    dataLogs.value.push(`[${new Date().toLocaleTimeString()}] ${res.message}`);
  } catch (error) {
    console.log(`Docker: ${error}`);
  }
}

async function removeContainer(id) {
  dataLogs.value.push(
    `[${new Date().toLocaleTimeString()}] Đang xoá container ${id}`
  );
  try {
    const formData = new FormData();
    formData.append("id", id);
    const res = await axios.post(`${API.REMOVE_CONTAINER_DOCKER}`, formData);
    if (res.success) {
      getListContainer();
    }
    dataLogs.value.push(`[${new Date().toLocaleTimeString()}] ${res.message}`);
  } catch (error) {
    console.log(`Docker: ${error}`);
  }
}
</script>

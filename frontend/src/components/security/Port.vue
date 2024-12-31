<template>
  <div class="m-2 bg-white border rounded-md p-9">
    <h1 class="mb-6 text-3xl font-bold text-gray-900">Cổng dịch vụ</h1>

    <div class="mb-6 overflow-hidden bg-white rounded-lg shadow-sm">
      <div class="p-4 border-b border-gray-200 sm:p-6">
        <h2 class="mb-4 text-xl font-semibold text-gray-900">Thêm cổng mới</h2>
        <form @submit.prevent="addPort" class="space-y-4">
          <div class="grid grid-cols-1 gap-4 sm:grid-cols-3">
            <div>
              <label for="port" class="block text-sm font-medium text-gray-700">
                Cổng
              </label>
              <input
                id="port"
                v-model="newPort.number"
                type="text"
                required
                class="w-full px-3 py-2 mt-1 border border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500"
              />
            </div>
            <div>
              <label
                for="port-protocol"
                class="block text-sm font-medium text-gray-700"
              >
                Giao thức
              </label>
              <select
                id="port-protocol"
                v-model="newPort.protocol"
                required
                class="block w-full py-3 pl-3 pr-10 mt-1 text-base border-gray-300 rounded-md focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
              >
                <option value="TCP">TCP</option>
                <option value="UDP">UDP</option>
              </select>
            </div>
            <div>
              <label
                for="port-status"
                class="block text-sm font-medium text-gray-700"
              >
                Trạng thái
              </label>
              <select
                id="port-status"
                v-model="newPort.status"
                required
                class="block w-full py-3 pl-3 pr-10 mt-1 text-base border-gray-300 rounded-md focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
              >
                <option value="ACCEPT">Cho phép</option>
                <option value="DROP">Chặn</option>
              </select>
            </div>
          </div>

          <div>
            <button
              type="submit"
              class="inline-flex items-center px-4 py-2 text-sm font-medium text-white bg-blue-600 border border-transparent rounded-md shadow-sm hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
            >
              <PlusIcon class="w-5 h-5 mr-2" />
              Thêm cổng
            </button>
          </div>
        </form>
      </div>

      <div class="p-4 sm:p-6">
        <h2 class="mb-4 text-xl font-semibold text-gray-900">Danh sách cổng</h2>
        <div class="overflow-x-auto">
          <table class="min-w-full border divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th
                  scope="col"
                  class="px-6 py-3 text-xs font-medium tracking-wider text-left text-gray-500 uppercase"
                >
                  Port
                </th>
                <th
                  scope="col"
                  class="px-6 py-3 text-xs font-medium tracking-wider text-left text-gray-500 uppercase"
                >
                  Protocol
                </th>

                <th
                  scope="col"
                  class="px-6 py-3 text-xs font-medium tracking-wider text-left text-gray-500 uppercase"
                >
                  Actions
                </th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr v-for="(port, i) in listPort" :key="`Port-ports-${i}`">
                <td
                  class="px-6 py-4 text-sm font-medium text-gray-900 whitespace-nowrap"
                >
                  {{ port.status == "DROP" ? "Chặn" : "Cho phép " }}
                  {{ port.port }}
                </td>
                <td class="px-6 py-4 text-sm text-gray-500 whitespace-nowrap">
                  {{ port.protocol }}
                </td>
                <td
                  class="flex px-6 py-4 text-sm font-medium text-right whitespace-nowrap"
                >
                  <button
                    @click="dropPort(port.port, port.protocol, port.status)"
                    class="text-red-600 hover:text-red-900"
                  >
                    <TrashIcon class="w-5 h-5" />
                    <span class="sr-only">Delete port {{ port.number }}</span>
                  </button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { PlusIcon, TrashIcon } from "lucide-vue-next";
import axios from "@/axios";
import API from "@/api";
import { useToast } from "vue-toastification";

const toast = useToast();
const listPort = ref([]);

const newPort = ref({
  number: "",
  protocol: "TCP",
  status: "ACCEPT",
});

getListPort();

async function getListPort() {
  try {
    const res = await axios.get(`${API.LIST_PORT}`);
    if (res.success) {
      listPort.value = res.message;
    } else {
      console.log(`Port: ${res.message}`);
    }
  } catch (error) {
    console.log(`Port: ${error}`);
  }
}

async function addPort() {
  try {
    const formData = new FormData();
    formData.append("port", newPort.value.number);
    formData.append("protocol", newPort.value.protocol);
    formData.append("status", newPort.value.status);
    const res = await axios.post(`${API.ADD_PORT}`, formData);
    if (res.success) {
      newPort.value.number = "";
      toast.success("Thêm cổng thành công");
      getListPort();
    } else {
      toast.error("Lỗi thêm cổng dịch vụ");
      console.log(`Port: ${res.message}`);
    }
  } catch (error) {
    console.log(`Port: ${error}`);
  }
}

async function dropPort(port, protocol, status) {
  try {
    const formData = new FormData();
    formData.append("port", port);
    formData.append("protocol", protocol);
    formData.append("status", status);
    const res = await axios.post(`${API.REMOVE_PORT}`, formData);
    if (res.success) {
      toast.success("Xóa cổng thành công");
      getListPort();
    } else {
      toast.error("Lỗi xóa cổng dịch vụ");
      console.log(`Port: ${res.message}`);
    }
  } catch (error) {
    console.log(`Port: ${error}`);
  }
}
</script>

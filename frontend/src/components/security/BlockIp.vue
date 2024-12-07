<template>
  <div class="m-2 bg-white border rounded-md p-9">
    <h1 class="mb-6 text-3xl font-bold text-gray-900">Chặn IP</h1>

    <div class="mb-6 overflow-hidden bg-white rounded-lg shadow-sm">
      <div class="p-4 border-b border-gray-200 sm:p-6">
        <h2 class="mb-4 text-xl font-semibold text-gray-900">
          Thêm địa chỉ IP chặn
        </h2>
        <form
          @submit.prevent="addBlockedIp"
          class="flex flex-col gap-4 sm:flex-row"
        >
          <div class="flex-grow">
            <label for="ip-address" class="sr-only">Địa chỉ IP</label>
            <input
              id="ip-address"
              v-model="newIpAddress"
              type="text"
              placeholder="Nhập địa chỉ IP"
              class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500"
            />
          </div>
          <button
            type="submit"
            class="inline-flex items-center px-4 py-2 text-sm font-medium text-white bg-blue-600 border border-transparent rounded-md shadow-sm hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          >
            <PlusIcon class="w-5 h-5 mr-2" />
            Thêm
          </button>
        </form>
      </div>

      <div class="p-4 sm:p-6">
        <h2 class="mb-4 text-xl font-semibold text-gray-900">
          Danh sách IP chặn
        </h2>
        <ul v-if="blockedIps.length" class="divide-y divide-gray-200">
          <li
            v-for="(ip, i) in blockedIps"
            :key="`BLockIp-blockedIps-${i}`"
            class="flex items-center justify-between py-4"
          >
            <span class="text-gray-900">{{ ip }}</span>
            <button
              @click="removeBlockedIp(ip)"
              class="ml-4 text-red-600 hover:text-red-800 focus:outline-none focus:underline"
            >
              <TrashIcon class="w-5 h-5" />
            </button>
          </li>
        </ul>
        <p v-else class="py-4 text-center text-gray-500">
          Hiện tại không có địa chỉ IP nào bị chặn.
        </p>
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
const blockedIps = ref([]);
const newIpAddress = ref("");

getBlockedIps();

async function getBlockedIps() {
  try {
    const res = await axios.get(`${API.LIST_BLOCK_IP}`);
    if (res.success) {
      blockedIps.value = [...res.message.trim().split("\n")].filter(
        (ip) => ip !== ""
      );
    } else {
      console.log(`BlockIp: ${res.message}`);
    }
  } catch (error) {
    console.log(`BlockIp: ${error}`);
  }
}

async function addBlockedIp() {
  try {
    const formData = new FormData();
    formData.append("ip", newIpAddress.value);
    const res = await axios.post(`${API.ADD_BLOCK_IP}`, formData);
    if (res.success) {
      newIpAddress.value = "";
      toast.success("Thêm ip thành công");
      getBlockedIps();
    } else {
      toast.error("Lỗi thêm ip");
      console.log(`BlockIp: ${res.message}`);
    }
  } catch (error) {
    console.log(`BlockIp: ${error}`);
  }
}

async function removeBlockedIp(ip) {
  try {
    const formData = new FormData();
    formData.append("ip", ip);
    const res = await axios.post(`${API.REMOVE_BLOCK_IP}`, formData);
    if (res.success) {
      toast.success("Xóa ip thành công");
      getBlockedIps();
    } else {
      toast.error("Lỗi xóa ip");
      console.log(`BlockIp: ${res.message}`);
    }
  } catch (error) {
    console.log(`BlockIp: ${error}`);
  }
}
</script>

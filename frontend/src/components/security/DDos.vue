<template>
  <div class="m-2 bg-white border rounded-md p-9">
    <h2 class="mb-4 text-xl font-semibold text-gray-900">Thêm quy tắc mới</h2>
    <form @submit.prevent="addRole" class="space-y-4">
      <div>
        <label for="new-rule" class="block text-sm font-medium text-gray-700">
          Câu lệnh
        </label>
        <input
          id="new-rule"
          type="text"
          v-model="role"
          required
          class="block w-full px-3 py-2 mt-1 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
          placeholder="iptables -A INPUT -p tcp --dport 80 -m conntrack --ctstate NEW -m recent --set --name rate_limit"
        />

        <ul class="mt-4 space-y-4">
          <li class="flex items-start">
            <code
              class="px-2 py-1 mr-3 font-mono text-sm text-gray-800 bg-gray-100 rounded"
            >
              --dport 80
            </code>
            <span class="text-gray-600">
              Áp dụng cho cổng 80 (HTTP). Bạn có thể thay đổi cổng nếu cần.
            </span>
          </li>

          <li class="flex items-start">
            <code
              class="w-[340px] px-2 py-1 mr-3 font-mono text-sm text-gray-800 bg-gray-100 rounded"
            >
              -m recent --set --name rate_limit
            </code>
            <span class="text-gray-600">
              Sử dụng module recent để đánh dấu địa chỉ IP của gói tin vào một
              danh sách có tên rate_limit. (cat /proc/net/xt_recent)
            </span>
          </li>
        </ul>
      </div>

      <div>
        <label for="new-rule" class="block text-sm font-medium text-gray-700">
          Câu lệnh
        </label>
        <input
          id="new-rule"
          type="text"
          value="iptables -A INPUT -p tcp --dport 80 -m conntrack --ctstate NEW -m recent --update --seconds 60 --hitcount 5 --name rate_limit -j DROP"
          required
          class="block w-full px-3 py-2 mt-1 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
          placeholder="iptables -A INPUT -p tcp --dport 80 -m conntrack --ctstate NEW -m recent --update --seconds 60 --hitcount 5 --name rate_limit -j DROP"
        />

        <ul class="mt-4 space-y-4">
          <li class="flex items-start">
            <code
              class="px-2 py-1 mr-3 font-mono text-sm text-gray-800 bg-gray-100 rounded"
            >
              --dport 80
            </code>
            <span class="text-gray-600">
              Áp dụng cho cổng 80 (HTTP). Bạn có thể thay đổi cổng nếu cần.
            </span>
          </li>

          <li class="flex items-start">
            <code
              class="px-2 py-1 mr-3 font-mono text-sm text-gray-800 bg-gray-100 rounded"
            >
              --seconds 60
            </code>
            <span class="text-gray-600">
              Giới hạn kiểm tra trong 60 giây gần đây.
            </span>
          </li>

          <li class="flex items-start">
            <code
              class="px-2 py-1 mr-3 font-mono text-sm text-gray-800 bg-gray-100 rounded"
            >
              --hitcount 5
            </code>
            <span class="text-gray-600">
              Nếu địa chỉ IP xuất hiện từ 5 lần trở lên trong vòng 60 giây, quy
              tắc này sẽ khớp.
            </span>
          </li>

          <li class="flex items-start">
            <code
              class="px-2 py-1 mr-3 font-mono text-sm text-gray-800 bg-gray-100 rounded"
            >
              -j DROP
            </code>
            <span class="text-gray-600">
              Nếu quy tắc khớp, gói tin sẽ bị DROP (loại bỏ).
            </span>
          </li>
        </ul>
      </div>

      <!-- <div>
        <button
          type="submit"
          class="inline-flex items-center px-4 py-2 text-sm font-medium text-white bg-blue-600 border border-transparent rounded-md shadow-sm hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
        >
          Thêm
        </button>
      </div> -->

      <div class="pt-5">
        <OpenTerminalButton />
      </div>
    </form>

    <!-- <div class="pt-6">
      <h2 class="mb-4 text-xl font-semibold">Logs</h2>
      <div
        class="h-48 p-4 overflow-y-auto font-mono text-sm text-green-400 bg-gray-900 rounded-lg"
      >
        <p v-for="(log, i) in logs" :key="`Mariadb-log-${i}`">{{ log }}</p>
      </div>
    </div> -->
  </div>
</template>

<script setup>
import { ref } from "vue";
import axios from "@/axios";
import API from "@/api";
import OpenTerminalButton from "@/components/common/OpenTerminalButton.vue";

const logs = ref([]);
const role = ref(
  "iptables -A INPUT -p tcp --dport 80 -m conntrack --ctstate NEW -m recent --set --name rate_limit"
);

async function addRole() {
  try {
    const formData = new FormData();
    formData.append("role", role.value);
    const res = await axios.post(`${API.ADD_ROLE}`, formData);

    if (res.success) {
      logs.value.push(
        `[${new Date().toLocaleTimeString()}] Đã thêm quy tắc thành công`
      );
    } else {
      logs.value.push(`[${new Date().toLocaleTimeString()}] ${res.message}`);
      console.log(`DDos: ${res.message}`);
    }
  } catch (error) {
    console.log(`DDos: ${error}`);
  }
}
</script>

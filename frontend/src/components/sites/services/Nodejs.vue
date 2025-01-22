<template>
  <div>
    <h3 class="pb-2">Nodejs</h3>
    <div class="mb-4">
      <label class="block mb-1 text-sm font-medium text-gray-700">
        Phiên bản
      </label>
      <select
        class="px-3 py-2 text-sm bg-transparent border border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500"
      >
        <option value="node:22">Node 22</option>
        <option value="node:20">Node 20</option>
        <option value="node:18">Node 18</option>
        <option value="node:16">Node 16</option>
        <option value="node:14">Node 14</option>
      </select>
    </div>

    <div class="mb-4">
      <label class="block mb-1 text-sm font-medium text-gray-700">
        Command
      </label>
      <input
        id="ip-address"
        type="text"
        placeholder="Nhập địa chỉ IP"
        value="npm i && npm run start"
        class="px-3 py-2 text-sm border border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500"
      />
    </div>
    <hr />
  </div>
</template>

<script setup>
import { reactive } from "vue";
const serviceConfig = reactive({
  nodejs: {
    container_name: "dddd",
    working_dir: "/usr/src/app",
    volumes: [".:/usr/src/app", "/usr/src/app/node_modules"],
    ports: ["3000:3000"],
    command: "npm run start",
  },
});
</script>

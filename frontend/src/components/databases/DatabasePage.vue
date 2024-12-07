<template>
  <div class="p-3 m-2 bg-white border rounded-md">
    <div class="flex items-center justify-between p-4 rounded-lg">
      <div class="flex items-center space-x-3">
        <h2 class="text-xl font-semibold text-gray-900">Cơ sở dữ liệu</h2>
      </div>
    </div>

    <div class="pl-4 mt-2 space-y-2">
      <RouterLink
        to="/databases/mariadb"
        class="flex items-center p-4 space-x-3 transition-colors rounded-lg hover:bg-blue-50 group"
      >
        <div
          class="flex items-center justify-center w-10 h-10 bg-blue-100 rounded-full"
        >
          <Database class="w-6 h-6 text-blue-600" />
        </div>
        <div>
          <span
            class="text-lg font-medium text-blue-700 group-hover:text-blue-800"
          >
            Mariadb
          </span>
          <p class="text-sm text-gray-600 hover:text-blue-700">
            <a
              href="https://mariadb.com"
              target="_blank"
              rel="noopener noreferrer"
            >
              https://mariadb.com
            </a>
          </p>
        </div>
      </RouterLink>
    </div>
  </div>
</template>

<script setup>
import { Database } from "lucide-vue-next";
</script>

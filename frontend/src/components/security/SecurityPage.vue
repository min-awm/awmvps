<template>
  <div class="p-3 m-2 bg-white border rounded-md">
    <!-- Header -->

    <div class="flex items-center p-6 space-x-2">
      <Shield class="w-6 h-6 text-gray-700" />
      <h2 class="text-lg font-medium text-gray-900">Bảo mật</h2>
    </div>

    <!-- Grid Content -->
    <div class="p-4">
      <div class="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-3">
        <!-- SSH Access -->
        <!-- <button
          class="flex items-center p-3 space-x-3 transition-colors rounded-lg hover:bg-gray-50"
        >
          <div
            class="flex items-center justify-center flex-shrink-0 w-10 h-10 rounded-lg bg-orange-50"
          >
            <Terminal class="w-5 h-5 text-orange-600" />
          </div>
          <span class="text-sm font-medium text-gray-900">SSH Access</span>
        </button> -->

        <!-- IP Blocker -->
        <RouterLink
          to="/security/block-ip"
          class="flex items-center p-3 space-x-3 transition-colors rounded-lg hover:bg-gray-50"
        >
          <div
            class="flex items-center justify-center flex-shrink-0 w-10 h-10 rounded-lg bg-blue-50"
          >
            <ShieldBan class="w-5 h-5 text-blue-600" />
          </div>
          <span class="text-sm font-medium text-gray-900">Chặn IP</span>
        </RouterLink>

        <!-- SSL/TLS -->
        <RouterLink
          to="/security/ssl"
          class="flex items-center p-3 space-x-3 transition-colors rounded-lg hover:bg-gray-50"
        >
          <div
            class="flex items-center justify-center flex-shrink-0 w-10 h-10 rounded-lg bg-green-50"
          >
            <Lock class="w-5 h-5 text-green-600" />
          </div>
          <span class="text-sm font-medium text-gray-900">
            Cấu hình SSL/TLS
          </span>
        </RouterLink>

        <!-- Activity -->
        <RouterLink
          to="/security/iftop"
          class="flex items-center p-3 space-x-3 transition-colors rounded-lg hover:bg-gray-50"
        >
          <div
            class="flex items-center justify-center flex-shrink-0 w-10 h-10 rounded-lg bg-purple-50"
          >
            <Activity class="w-5 h-5 text-purple-600" />
          </div>
          <span class="text-sm font-medium text-gray-900">
            Theo dõi truy cập
          </span>
        </RouterLink>

        <!-- DatabaseBackup -->
        <!-- <button
          class="flex items-center p-3 space-x-3 transition-colors rounded-lg hover:bg-gray-50"
        >
          <div
            class="flex items-center justify-center flex-shrink-0 w-10 h-10 rounded-lg bg-red-50"
          >
            <DatabaseBackup class="w-5 h-5 text-red-600" />
          </div>
          <span class="text-sm font-medium text-gray-900">Backup</span>
        </button> -->

        <!-- Leech Protection -->
        <RouterLink
          to="/security/htop"
          class="flex items-center p-3 space-x-3 transition-colors rounded-lg hover:bg-gray-50"
        >
          <div
            class="flex items-center justify-center flex-shrink-0 w-10 h-10 rounded-lg bg-yellow-50"
          >
            <Cpu class="w-5 h-5 text-yellow-600" />
          </div>
          <span class="text-sm font-medium text-gray-900">
            Theo dõi tiến trình
          </span>
        </RouterLink>

        <!-- SSL/TLS Status -->
        <!-- <button
          class="flex items-center p-3 space-x-3 transition-colors rounded-lg hover:bg-gray-50"
        >
          <div
            class="flex items-center justify-center flex-shrink-0 w-10 h-10 rounded-lg bg-cyan-50"
          >
            <FileCheck class="w-5 h-5 text-cyan-600" />
          </div>
          <span class="text-sm font-medium text-gray-900">SSL/TLS Status</span>
        </button> -->

        <!-- Two-Factor Authentication -->
        <RouterLink
          to="/security/port"
          class="flex items-center p-3 space-x-3 transition-colors rounded-lg hover:bg-gray-50"
        >
          <div
            class="flex items-center justify-center flex-shrink-0 w-10 h-10 rounded-lg bg-indigo-50"
          >
            <Key class="w-5 h-5 text-indigo-600" />
          </div>
          <span class="text-sm font-medium text-gray-900">Cổng dịch vụ</span>
        </RouterLink>

        <!-- Imunify360 -->
        <RouterLink
          to="/security/ddos"
          class="flex items-center p-3 space-x-3 transition-colors rounded-lg hover:bg-gray-50"
        >
          <div
            class="flex items-center justify-center flex-shrink-0 w-10 h-10 rounded-lg bg-rose-50"
          >
            <ShieldCheck class="w-5 h-5 text-rose-600" />
          </div>
          <span class="text-sm font-medium text-gray-900">Chống DDOS</span>
        </RouterLink>
      </div>
    </div>
  </div>
</template>

<script setup>
import {
  Shield,
  Lock,
  Activity,
  Key,
  ShieldCheck,
  ShieldBan,
  Cpu,
} from "lucide-vue-next";
import { RouterLink } from "vue-router";
</script>

<template>
  <div class="m-5 space-y-6">
    <!-- Feature Section -->
    <div class="bg-white rounded-lg shadow">
      <button
        @click="showFeature = !showFeature"
        class="flex items-center justify-between w-full px-4 py-3 text-gray-700 hover:bg-gray-50"
      >
        <div class="flex items-center space-x-2">
          <Feather class="w-5 h-5 text-gray-500" />
          <span class="text-lg font-medium">Chức năng</span>
        </div>
        <ChevronDown
          class="w-5 h-5 text-gray-400 transition-transform duration-200"
          :class="{ 'rotate-180': showFeature }"
        />
      </button>

      <div v-if="showFeature" class="p-4">
        <div class="grid grid-cols-2 gap-4 sm:grid-cols-3 md:grid-cols-5">
          <RouterLink
            v-for="(item, i) in featureItems"
            :to="item.link"
            :key="`HomePage-featureItems-${i}`"
            class="flex flex-col items-center p-4 transition-colors rounded-lg hover:bg-gray-50"
          >
            <component :is="item.icon" class="w-8 h-8 text-blue-500" />
            <span class="mt-2 text-sm text-center text-gray-600">
              {{ item.name }}
            </span>
          </RouterLink>
        </div>
      </div>
    </div>

    <!-- Database Section -->
    <div class="bg-white rounded-lg shadow">
      <button
        @click="showDatabase = !showDatabase"
        class="flex items-center justify-between w-full px-4 py-3 text-gray-700 hover:bg-gray-50"
      >
        <div class="flex items-center space-x-2">
          <Database class="w-5 h-5 text-gray-500" />
          <span class="text-lg font-medium">Cơ sở dữ liệu</span>
        </div>
        <ChevronDown
          class="w-5 h-5 text-gray-400 transition-transform duration-200"
          :class="{ 'rotate-180': showDatabase }"
        />
      </button>

      <div v-if="showDatabase" class="p-4">
        <div class="grid grid-cols-2 gap-4 sm:grid-cols-3 md:grid-cols-4">
          <RouterLink
            v-for="(item, i) in databaseItems"
            :to="item.link"
            :key="`HomePage-databaseItems-${i}`"
            class="flex flex-col items-center p-4 transition-colors rounded-lg hover:bg-gray-50"
          >
            <component :is="item.icon" class="w-8 h-8 text-blue-500" />
            <span class="mt-2 text-sm text-center text-gray-600">
              {{ item.name }}
            </span>
          </RouterLink>
        </div>
      </div>
    </div>

    <!-- Security Section -->
    <div class="bg-white rounded-lg shadow">
      <button
        @click="showSecurity = !showSecurity"
        class="flex items-center justify-between w-full px-4 py-3 text-gray-700 hover:bg-gray-50"
      >
        <div class="flex items-center space-x-2">
          <Shield class="w-5 h-5 text-gray-500" />
          <span class="text-lg font-medium">Cơ sở dữ liệu</span>
        </div>
        <ChevronDown
          class="w-5 h-5 text-gray-400 transition-transform duration-200"
          :class="{ 'rotate-180': showSecurity }"
        />
      </button>

      <div v-if="showSecurity" class="p-4">
        <div class="grid grid-cols-2 gap-4 sm:grid-cols-3 md:grid-cols-4">
          <RouterLink
            v-for="(item, i) in securityItems"
            :to="item.link"
            :key="`HomePage-securityItems-${i}`"
            class="flex flex-col items-center p-4 transition-colors rounded-lg hover:bg-gray-50"
          >
            <component :is="item.icon" class="w-8 h-8 text-blue-500" />
            <span class="mt-2 text-sm text-center text-gray-600">
              {{ item.name }}
            </span>
          </RouterLink>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import {
  Feather,
  Database,
  ChevronDown,
  Folder,
  Puzzle,
  Globe,
  Terminal,
  Shield,
  ShieldCheck,
  Key,
  Activity,
  Lock,
  ShieldBan,
  Cpu,
} from "lucide-vue-next";

const showFeature = ref(true);
const showDatabase = ref(true);
const showSecurity = ref(true);

const featureItems = [
  { name: "Quản lý tập tin", icon: Folder, link: "/files" },
  { name: "Docker", icon: Globe, badge: Puzzle, link: "/docker" },
  { name: "Nginx", icon: Globe, badge: Globe, link: "/nginx" },
  { name: "Terminal", icon: Terminal, link: "/terminal" },
];

const databaseItems = [
  { name: "Mariadb", icon: Database, link: "/databases/mariadb" },
];

const securityItems = [
  { name: "Chặn IP", icon: ShieldBan, link: "/security/block-ip" },
  { name: "Cấu hình SSL/TLS", icon: Lock, link: "/security/ssl" },
  { name: "Theo dõi truy cập", icon: Activity, link: "/security/iftop" },
  { name: "Theo dõi tiến trình", icon: Cpu, link: "/security/htop" },
  { name: "Cổng dịch vụ", icon: Key, link: "/security/port" },
  { name: "Chống DDOS", icon: ShieldCheck, link: "/security/ddos" },
];
</script>

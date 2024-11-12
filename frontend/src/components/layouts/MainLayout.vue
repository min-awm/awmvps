<template>
  <div class="flex flex-col h-screen bg-gray-100">
    <!-- Header -->
    <header class="flex items-center justify-between p-4 bg-white shadow-sm">
      <div class="flex items-center">
        <button
          @click="toggleSidebar"
          class="p-2 mr-2 rounded-md md:hidden hover:bg-gray-100"
        >
          <AlignJustify class="w-6 h-6" />
        </button>
        <h1 class="text-xl font-bold text-blue-600">NAMEAPP</h1>
      </div>
      <div class="flex items-center space-x-4">
        <Bell class="w-6 h-6 text-gray-500" />
        <div class="flex items-center space-x-2">
          <div
            class="flex items-center justify-center w-8 h-8 text-gray-600 bg-gray-300 rounded-full"
          >
            TM
          </div>
          <div class="hidden md:block">
            <div class="font-medium">Trần Nhật Minh</div>
            <div class="text-sm text-gray-500">minha0512003a@gmail.com</div>
          </div>
          <ChevronDown class="w-4 h-4 text-gray-500" />
        </div>
      </div>
    </header>

    <div class="flex flex-1">
      <!-- Sidebar -->
      <aside
        :class="[
          'bg-white w-64 p-4 shadow-md',
          { hidden: !isSidebarOpen },
          'md:block',
        ]"
      >
        <nav class="space-y-2">
          <RouterLink
            v-for="(item, i) in menuItems"
            :key="`MainLayout-menuItems-${i}`"
            :to="item.route"
            class="flex items-center w-full px-4 py-2 text-left rounded-md hover:bg-gray-100 menu-item"
          >
            <component :is="item.icon" class="w-4 h-4 mr-2" />
            {{ item.text }}
          </RouterLink>
        </nav>
      </aside>

      <router-view />
    </div>

    <footer class="p-4 text-sm text-center text-gray-600 bg-white">
      Copyright © 2024 NAMEAPP. All Rights Reserved.
    </footer>
  </div>
</template>

<script setup>
import { ref } from "vue";
import {
  Bell,
  ChevronDown,
  Home,
  Folder,
  Globe,
  Database,
  ChartNoAxesCombined,
  Puzzle,
  Shield,
  HelpCircle,
  Phone,
  AlignJustify,
  Terminal,
} from "lucide-vue-next";

const isSidebarOpen = ref(false);

const toggleSidebar = () => {
  isSidebarOpen.value = !isSidebarOpen.value;
};

const menuItems = [
  { icon: Home, text: "Trang chủ", route: "/" },
  { icon: Folder, text: "Tập tin", route: "/files" },
  { icon: Puzzle, text: "Môi trường", route: "/environment" },
  { icon: Globe, text: "Quản lý tên miền", route: "/domains" },
  { icon: Database, text: "Cơ sở dữ liệu", route: "/databases" },
  { icon: Terminal, text: "Terminal", route: "/terminal" },
  { icon: Shield, text: "Bảo mật", route: "/security" },
  { icon: ChartNoAxesCombined, text: "Thống kê", route: "/statistic" },
  { icon: HelpCircle, text: "Hướng dẫn", route: "/tutorial" },
  { icon: Phone, text: "Hỗ trợ: 1234 5678", route: "/support" },
];
</script>

<style scoped lang="scss">
.menu-item {
  &.router-link-exact-active {
    background-color: #38a7ff;
    color: #fff;
  }
}
</style>

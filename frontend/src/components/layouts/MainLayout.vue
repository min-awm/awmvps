<template>
  <div class="relative flex flex-col min-h-screen bg-gray-100">
    <!-- Header -->
    <header
      class="flex items-center justify-between p-4 bg-white shadow-sm border-b h-[64px]"
    >
      <div class="flex items-center">
        <button
          @click="toggleSidebar"
          class="p-2 mr-2 rounded-md md:hidden hover:bg-gray-100"
        >
          <AlignJustify class="w-6 h-6" />
        </button>
        <h1 class="text-xl font-bold text-blue-600">AWMVPS</h1>
      </div>
      <div class="flex items-center space-x-4">
        <Bell class="w-6 h-6 text-gray-500" />
        <div class="relative z-100">
          <div
            class="flex items-center space-x-2 cursor-pointer"
            @click="isUserDropDownOpen = !isUserDropDownOpen"
          >
            <div
              class="flex items-center justify-center w-8 h-8 text-gray-600 bg-gray-300 rounded-full"
            >
              {{ userInfo?.name ? userInfo.name[0] : "" }}
            </div>
            <div class="hidden md:block">
              <div class="font-medium">{{ userInfo?.name }}</div>
              <div class="text-sm text-gray-500"></div>
            </div>
            <ChevronDown
              :class="[
                'w-4 h-4 text-gray-500  transition',
                isUserDropDownOpen ? 'rotate-180' : 'rotate-0',
              ]"
            />
          </div>

          <div
            v-if="isUserDropDownOpen"
            class="absolute right-0 z-50 w-56 py-1 mt-2 bg-white border border-gray-100 rounded-lg shadow-lg top-8"
          >
            <div class="px-4 py-2 border-b border-gray-100">
              <div class="text-sm font-medium text-gray-900">
                {{ userInfo?.name }}
              </div>
              <div class="text-xs text-gray-500">Người quản lý</div>
            </div>

            <div class="py-1">
              <div
                v-for="(item, i) in userItems"
                :key="`MainLayout-userItems-${i}`"
                class="flex items-center px-4 py-2 text-sm text-gray-700 transition-colors duration-150 cursor-pointer hover:bg-blue-50 hover:text-blue-600"
                @click="handleUserItems(item.key)"
              >
                <component :is="item.icon" class="w-5 h-5 mr-3" />
                {{ item.label }}
              </div>
            </div>
          </div>
        </div>
      </div>
    </header>

    <div class="flex flex-1">
      <!-- Sidebar -->
      <aside
        :class="[
          'bg-white w-[190px] p-4 border-r absolute bottom-0 top-[64px] min-h-[500px] z-10',
          { hidden: !isSidebarOpen },
          'md:block md:relative md:top-0',
        ]"
      >
        <nav class="space-y-2">
          <RouterLink
            v-for="(item, i) in menuItems"
            :key="`MainLayout-menuItems-${i}`"
            class="flex items-center px-4 py-2 text-left rounded-md hover:bg-gray-100 menu-item"
            :to="item.route"
            @click="isSidebarOpen = false;"
          >
            <component :is="item.icon" class="w-4 h-4 mr-2" />
            {{ item.text }}
          </RouterLink>
        </nav>
      </aside>

      <div class="w-screen md:w-[calc(100vw-190px)] min-h-[500px]">
        <router-view />
      </div>
    </div>
  </div>

  <PasswordChangeModal
    v-if="isPasswordChangeModalOpen"
    @close="isPasswordChangeModalOpen = false"
  />
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
  Puzzle,
  Shield,
  HelpCircle,
  Phone,
  AlignJustify,
  Terminal,
  Lock,
  LogOut,
} from "lucide-vue-next";
import { useUserStore } from "@/store/user";
import { storeToRefs } from "pinia";
import PasswordChangeModal from "@/components/layouts/PasswordChangeModal.vue";

const isSidebarOpen = ref(false);
const userStore = useUserStore();
const { userInfo } = storeToRefs(userStore);
const userItems = [
  {
    key: "changePassword",
    label: "Mật khẩu",
    icon: Lock,
  },
  {
    key: "logout",
    label: "Đăng xuất",
    icon: LogOut,
  },
];
const isUserDropDownOpen = ref(false);
const isPasswordChangeModalOpen = ref(false);

const menuItems = [
  { icon: Home, text: "Trang chủ", route: "/home" },
  { icon: Folder, text: "Tập tin", route: "/files" },
  { icon: Puzzle, text: "Docker", route: "/docker" },
  { icon: Globe, text: "Nginx", route: "/nginx" },
  { icon: Database, text: "Cơ sở dữ liệu", route: "/databases" },
  { icon: Terminal, text: "Terminal", route: "/terminal" },
  { icon: Shield, text: "Bảo mật", route: "/security" },
  { icon: HelpCircle, text: "Hướng dẫn", route: "/tutorial" },
  { icon: Phone, text: "1234 5678", route: "/support" },
];

function toggleSidebar() {
  isSidebarOpen.value = !isSidebarOpen.value;
}

function handleUserItems(key) {
  switch (key) {
    case "logout":
      userStore.logout();
      break;
    case "changePassword":
      isPasswordChangeModalOpen.value = true;
      break;

    default:
      break;
  }
}
</script>

<style scoped lang="scss">
.menu-item {
  &.router-link-active {
    background-color: #38a7ff;
    color: #fff;
  }
}
</style>

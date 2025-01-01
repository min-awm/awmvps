<template>
  <div class="m-2 bg-white border rounded-md">
    <!-- Top Navigation -->
    <header class="border-b">
      <div class="flex items-center justify-between px-4 h-14">
        <div class="flex items-center justify-between w-full space-x-2">
          <!-- Mobile Menu Toggle -->
          <button
            @click="isSidebarOpen = !isSidebarOpen"
            class="p-2 rounded-md hover:bg-gray-100"
          >
            <Bookmark class="w-5 h-5 text-gray-600" />
          </button>

          <div
            class="flex items-center space-x-2 px-3 py-2 overflow-x-auto scroll-bar bg-gray-100 rounded-md w-[60%]"
          >
            <ChevronLeft @click="backPath()" class="cursor-pointer" />
            <input
              type="text"
              :value="path"
              class="w-full bg-transparent outline-none"
              @input="(e) => (pathTemp = e.target.value)"
              @focus="showSendPath = true"
              @blur="sendPathFn()"
              @keyup.enter="goPath()"
              placeholder="Nhập đường dẫn"
            />
            <Forward v-if="showSendPath" @click="goPath()" />
            <RefreshCw
              size="18"
              :class="`cursor-pointer  ${
                refreshDirectory ? 'animate-spin' : ''
              }`"
              @click="refreshDirectoryFn()"
              v-else
            />
          </div>

          <div class="flex items-center space-x-2">
            <UploadFile :path="path" @uploadDone="refreshDirectoryFn()" />
          </div>
        </div>
      </div>
    </header>

    <div class="flex">
      <!-- Sidebar -->
      <div
        class="fixed top-0 bottom-0 left-0 right-0 md:hidden"
        v-if="isSidebarOpen"
        @click="isSidebarOpen = false"
      ></div>
      <aside
        :class="`${
          isSidebarOpen ? 'translate-x-0 ' : '-translate-x-full hidden'
        } scroll-bar overflow-y-auto h-screen md:h-[calc(100vh-140px)] md:translate-x-0 fixed md:static inset-y-0 left-0 w-64 border-r transform bg-white transition-transform duration-200 ease-in-out z-30`"
      >
        <!-- <div
          class="text-sm mx-6 border-t pt-2 flex gap-2 items-center absolute top-[60px] right-0 left-0"
          v-if="sidebarItems.length === 1"
        >
          <Plus size="18" />
          Dấu trang
        </div> -->
        <nav class="p-4 space-y-1">
          <draggable
            class="md:h-[calc(100vh-180px)] h-screen"
            :list="sidebarItems"
            group="forders"
            @change="changeSideBarItems"
            item-key="name"
          >
            <template #item="{ element }">
              <div
                :class="`flex items-center space-x-3 px-3 py-2 rounded-md text-sm cursor-pointer mb-2 ${
                  element?.active ? 'bg-gray-100' : 'hover:bg-gray-100'
                }`"
                @click="checkPath('')"
              >
                <component :is="element.icon" class="w-5 h-5 text-gray-600" />
                <span>{{ element.name }}</span>
              </div>
            </template>
          </draggable>
        </nav>
      </aside>

      <!-- Main Content -->
      <main class="flex-1 p-6 overflow-y-auto h-[calc(100vh-140px)] scroll-bar">
        <draggable
          class="grid grid-cols-2 gap-4 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-6"
          :list="folders"
          :group="{ name: 'forders', pull: 'clone', put: false }"
          item-key="fileManager-forders"
          :sort="false"
        >
          <template #item="{ element }">
            <div
              class="flex flex-col items-center p-4 transition-colors rounded-lg cursor-pointer hover:bg-gray-100"
              @click="openFileForder(element.type, element.path)"
            >
              <div
                class="flex items-center justify-center w-16 h-16 mb-2 bg-blue-100 rounded-lg"
              >
                <component :is="element.icon" class="w-8 h-8 text-blue-600" />
              </div>
              <span class="text-sm text-center text-gray-700">
                {{ element.name }}
              </span>
            </div>
          </template>
        </draggable>
      </main>
    </div>
  </div>

  <!-- <RenameModal /> -->
</template>

<script setup>
import { ref, onMounted } from "vue";
import {
  Home,
  File,
  Folder,
  Plus,
  Bookmark,
  RefreshCw,
  Forward,
  ChevronLeft,
  AppWindowMac,
  TriangleAlert,
} from "lucide-vue-next";
import draggable from "vuedraggable";
import axios from "@/axios";
import API from "@/api";
import { useToast } from "vue-toastification";
import UploadFile from "@/components/files/UploadFile.vue";
// import RenameModal from "@/components/files/RenameModal.vue";

const toast = useToast();
const isSidebarOpen = ref(true);
const sidebarItems = ref([
  { name: "Home", icon: Home, active: true, path: "home" },
]);
const folders = ref();
const refreshDirectory = ref(false);
const showSendPath = ref(false);
const path = ref();
const pathTemp = ref("");

function changeSideBarItems(evt) {
  window.console.log(evt.added);
  return false;
}

async function checkPath(queryPath) {
  try {
    const res = await axios.get(`${API.CHECK_PATH}?path=${queryPath}`);
    if (res.success) {
      path.value = res.path;
      getList(res.path);
    } else {
      toast.error("Đường dẫn không tồn tại");
    }
  } catch (error) {
    console.log(`Filemanager: ${error}`);
    toast.error("Lỗi");
  }
}

async function getList(queryPath) {
  try {
    const res = await axios.get(`${API.GET_LIST}?path=${queryPath}`);
    if (res.success) {
      folders.value = res.list.map((item) => {
        let icon;
        switch (item.type) {
          case "directory":
            icon = Folder;
            break;
          case "file":
            icon = File;
            break;
          case "application":
            icon = AppWindowMac;
            break;
          default:
            icon = TriangleAlert;
        }

        return {
          ...item,
          icon,
          path: `${path.value}/${item.name}`,
        };
      });
    } else {
      switch (res.message) {
        case "cmdErr":
          toast.error("Không lấy được danh sách");
          break;

        default:
          break;
      }
    }
  } catch (error) {
    console.log(`Filemanager: ${error}`);
    toast.error("Lỗi");
  }
}

function goPath() {
  checkPath(pathTemp.value);
}

function refreshDirectoryFn() {
  refreshDirectory.value = true;
  setTimeout(() => {
    refreshDirectory.value = false;
  }, 500);
  checkPath(path.value);
}

function sendPathFn() {
  setTimeout(() => {
    showSendPath.value = false;
  }, 500);
}

function openFileForder(type, path) {
  switch (type) {
    case "directory":
      checkPath(path);
      break;
    case "file":
      break;
    case "application":
      toast.warning("Đây là file ứng dụng");
      break;

    default:
      break;
  }
}

function backPath() {
  let pathQuery = "/";
  const index = path.value.lastIndexOf("/");
  if (index > 0) {
    pathQuery = path.value.slice(0, index);
  }

  checkPath(pathQuery);
}

// async function rename() {
//   try {
//     const formData = new FormData();
//     formData.append("oldName", "/home/min/my-code/test/a1.html");
//     formData.append("newName", "/home/min/my-code/test/a6.html");

//     const res = await axios.post(API.RENAME_FILE, formData);
//     if (res.success) {
//       toast.success("Đổi tên thành công");
//       checkPath(path.value);
//     } else {
//       toast.error("Đổi tên lỗi");
//     }
//   } catch (error) {
//     console.log(`Filemanager: ${error}`);
//     toast.error("Lỗi");
//   }
// }

checkPath("");
onMounted(() => {
  if (window.innerWidth < 768) {
    isSidebarOpen.value = false;
  }
});
</script>

<style scoped lang="scss">
.scroll-bar {
  &::-webkit-scrollbar {
    width: 6px;
    height: 4px;
  }

  &::-webkit-scrollbar-track {
    background: #f1f1f1;
  }

  &::-webkit-scrollbar-thumb {
    background: #cecece;
    border-radius: 4px;
  }

  &::-webkit-scrollbar-thumb:hover {
    background: #bcbcbc;
  }
}
</style>

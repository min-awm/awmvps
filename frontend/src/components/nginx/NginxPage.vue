<template>
  <div class="p-3 m-2 bg-white border rounded-md">
    <div class="p-6" v-if="step == 0">
      <h2 class="mb-4 text-2xl font-semibold">Nginx</h2>

      <div class="flex items-center space-x-4">
        <div class="w-4 h-4 bg-gray-500 rounded-full"></div>
        <span class="text-lg">Chưa cài đặt</span>
      </div>
      <button
        @click="installNginx"
        class="mt-4 bg-green-500 btn-action hover:bg-green-600"
        :disabled="isProcessingInstall"
      >
        <Download v-if="!isProcessingInstall" class="w-5 h-5 mr-2" />
        <Loader2 v-else class="w-5 h-5 mr-2 animate-spin" />
        Cài đặt
      </button>
    </div>

    <div class="p-6" v-if="step == 1">
      <h2 class="mb-4 text-2xl font-semibold">Nginx</h2>

      <div class="flex items-center mb-3 space-x-4">
        <div :class="`w-4 h-4 ${colorStatus(status)} rounded-full`"></div>
        <span class="text-lg">{{ textStatus(status) }}</span>
        <button
          :class="` btn-action cursor-pointer ${
            status === 'active\n' ? 'bg-red-500' : 'bg-green-500'
          }`"
          :disabled="isProcessingStopStart"
          @click="stopStartNginx(status)"
        >
          <Loader2
            class="w-5 h-5 mr-2 animate-spin"
            v-if="isProcessingStopStart"
          />
          {{ status === "active\n" ? "Dừng" : "Chạy" }}
        </button>
      </div>

      <div class="overflow-hidden bg-white border rounded-lg shadow-sm">
        <div class="p-4 border-b border-gray-200 sm:p-6">
          <div class="sm:flex sm:items-center sm:justify-between">
            <div class="flex-1 min-w-0">
              <h2
                class="text-2xl font-bold leading-7 text-gray-900 sm:truncate"
              >
                {{ selectedConfig?.name }}
              </h2>
            </div>
            <div class="flex gap-4 mt-4 sm:mt-0 sm:ml-4">
              <button
                @click="removeFileConf"
                class="inline-flex items-center px-4 py-2 text-sm font-medium text-white bg-red-500 border border-transparent rounded-md shadow-sm hover:bg-red-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500"
              >
                Xóa file
              </button>
              <button
                @click="saveFileConf"
                class="inline-flex items-center px-4 py-2 text-sm font-medium text-white bg-blue-500 border border-transparent rounded-md shadow-sm hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
              >
                Lưu file
              </button>
            </div>
          </div>
        </div>

        <div class="flex flex-col sm:flex-row">
          <!-- Config List -->
          <div
            class="w-full p-4 border-b border-gray-200 sm:w-64 bg-gray-50 sm:p-6 sm:border-b-0 sm:border-r"
          >
            <h3 class="mb-4 text-lg font-medium text-gray-900">
              Tập tin cấu hình
            </h3>
            <input
              type="text"
              v-model="nameFileConf"
              class="w-full px-3 py-2 text-sm border border-gray-300 rounded-md shadow-sm outline-none"
              placeholder="Nhập tên file"
            />
            <button
              @click="createFileConf"
              class="flex items-center px-4 py-2 my-4 text-sm font-medium text-white bg-blue-500 border border-transparent rounded-md shadow-sm hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
            >
              Tạo file
            </button>
            <ul class="space-y-2">
              <li
                v-for="(config, i) in configs"
                :key="`NginxPage-configs-${i}`"
                @click="getFileConf(config)"
                class="p-2 transition-colors duration-150 rounded-md cursor-pointer hover:bg-gray-100"
                :class="{ 'bg-blue-100': config.id === selectedConfig?.id }"
              >
                {{ config.name }}
              </li>
            </ul>
          </div>

          <!-- Config Editor -->
          <div class="flex-1 p-4 sm:p-6">
            <div class="mb-4">
              <label
                for="config-content"
                class="block text-sm font-medium text-gray-700"
              >
                Edit Configuration
              </label>
              <textarea
                id="config-content"
                rows="20"
                class="block w-full mt-1 font-mono border border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
                v-model="contentFileConf"
              ></textarea>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="p-6 pt-0">
      <button
        @click="checkConf"
        class="flex items-center px-4 py-2 my-4 font-medium text-white bg-green-500 border border-transparent rounded-md shadow-sm hover:bg-green-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
      >
        Kiểm tra cấu hình
      </button>
      <p class="text-gray-600">Để cấu hình áp dụng, vui lòng tắt bật Nginx</p>
    </div>

    <div class="p-6 pt-0">
      <h2 class="mb-4 text-2xl font-semibold">Logs</h2>
      <div
        class="h-48 p-4 overflow-y-auto font-mono text-sm text-green-400 bg-gray-900 rounded-lg"
      >
        <p v-for="(log, i) in dataLogs" :key="`Nginx-log-${i}`">{{ log }}</p>
      </div>
    </div>

    <div class="p-6 py-0" v-if="step == 1">
      <button
        @click="removeNginx()"
        class="bg-red-500 btn-action hover:bg-red-600"
        :disabled="isProcessingRemove"
      >
        <Trash2 v-if="!isProcessingRemove" class="w-5 h-5 mr-2" />
        <Loader2 v-else class="w-5 h-5 mr-2 animate-spin" />
        Gỡ cài đặt
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { Download, Loader2, Trash2 } from "lucide-vue-next";
import axios from "@/axios";
import API from "@/api";

const step = ref(0);
const dataLogs = ref([]);
const status = ref("");
const isProcessingInstall = ref(false);
const isProcessingStopStart = ref(false);
const isProcessingRemove = ref(false);
const configs = ref([]);
const nameFileConf = ref("");
const contentFileConf = ref("");
const selectedConfig = ref(configs.value[0]);

checkStatus();
getListConf();

async function checkStatus() {
  const installingNginx = localStorage.getItem("installingNginx");
  if (installingNginx == "true") {
    isProcessingInstall.value = true;
  }
  const removingNginx = localStorage.getItem("removingNginx");

  if (removingNginx == "true") {
    isProcessingRemove.value = true;
  }

  try {
    const res = await axios.get(`${API.STATUS_NGINX}`);
    if (res.success) {
      status.value = res.status;
      step.value = 1;
      localStorage.setItem("installingNginx", false);
    } else {
      switch (res.message) {
        case "notInstall":
          step.value = 0;
          localStorage.setItem("removingNginx", false);
          break;

        default:
          break;
      }
    }
  } catch (error) {
    console.log(`Nginx: ${error}`);
  }
}

async function installNginx() {
  isProcessingInstall.value = true;
  dataLogs.value.push(
    `[${new Date().toLocaleTimeString()}] Quá trình cài đặt có thể mất vài phút.`
  );
  localStorage.setItem("installingNginx", true);
  try {
    const res = await axios.get(`${API.INSTALL_NGINX}`);
    if (res.success) {
      checkStatus();
    }
    dataLogs.value.push(`[${new Date().toLocaleTimeString()}] ${res.message}`);
  } catch (error) {
    console.log(`Nginx: ${error}`);
  } finally {
    isProcessingInstall.value = false;
    localStorage.setItem("installingNginx", false);
  }
}

function colorStatus(status) {
  let colorReturn = "";
  switch (status) {
    case "active\n":
      colorReturn = "bg-green-500";
      break;

    case "inactive\n":
      colorReturn = "bg-gray-500";
      break;

    default:
      colorReturn = "bg-red-500";
      break;
  }

  return colorReturn;
}

function textStatus(status) {
  let textReturn = "";
  switch (status) {
    case "active\n":
      textReturn = "Đang chạy";
      break;

    case "inactive\n":
      textReturn = "Đã dừng";
      break;

    default:
      textReturn = "Đang lỗi";
      break;
  }

  return textReturn;
}

async function stopStartNginx(status) {
  isProcessingStopStart.value = true;
  let url = API.START_NGINX;
  if (status == "active\n") {
    url = API.STOP_NGINX;
  }

  try {
    const res = await axios.get(`${url}`);
    if (res.success) {
      checkStatus();
    }

    dataLogs.value.push(`[${new Date().toLocaleTimeString()}] ${res.message}`);
  } catch (error) {
    console.log(`Nginx: ${error}`);
  } finally {
    isProcessingStopStart.value = false;
  }
}

async function removeNginx() {
  isProcessingRemove.value = true;
  dataLogs.value.push(
    `[${new Date().toLocaleTimeString()}] Quá trình gỡ cài đặt có thể mất vài phút.`
  );
  localStorage.setItem("removingNginx", true);

  try {
    const res = await axios.get(`${API.REMOVE_NGINX}`);
    if (res.success) {
      checkStatus();
    }

    dataLogs.value.push(`[${new Date().toLocaleTimeString()}] ${res.message}`);
  } catch (error) {
    console.log(`Nginx: ${error}`);
  } finally {
    isProcessingRemove.value = false;
    localStorage.setItem("removingNginx", false);
  }
}

async function getListConf() {
  try {
    const res = await axios.get(`${API.LIST_CONF_NGINX}`);
    if (res.success) {
      const fileList = res.message
        .trim()
        .split("\n")
        .map((name, index) => ({ id: index + 1, name }));
      if (fileList.length) {
        selectedConfig.value = fileList[0];
        getFileConf(fileList[0]);
      }
      configs.value = fileList;
    } else {
      if (
        res.message == "ls: cannot access '*.conf': No such file or directory\n"
      ) {
        configs.value = [];
        selectedConfig.value = null;
        contentFileConf.value = "";
      } else {
        dataLogs.value.push(
          `[${new Date().toLocaleTimeString()}] ${res.message}`
        );
      }
    }
  } catch (error) {
    console.log(`Nginx: ${error}`);
  }
}

async function getFileConf(config) {
  try {
    const formData = new FormData();
    formData.append("fileName", `${config.name}`);
    const res = await axios.post(API.FILE_CONF_NGINX, formData);
    if (res.success) {
      selectedConfig.value = config;
      contentFileConf.value = res.message;
    } else {
      dataLogs.value.push(
        `[${new Date().toLocaleTimeString()}] ${res.message}`
      );
    }
  } catch (error) {
    console.log(`Nginx: ${error}`);
  }
}

async function createFileConf() {
  try {
    const formData = new FormData();
    formData.append("fileName", `${nameFileConf.value}.conf`);
    const res = await axios.post(API.CREATE_CONF_NGINX, formData);
    if (res.success) {
      getListConf();
      nameFileConf.value = "";
    }

    dataLogs.value.push(`[${new Date().toLocaleTimeString()}] ${res.message}`);
  } catch (error) {
    console.log(`Nginx: ${error}`);
  }
}

async function removeFileConf() {
  try {
    if (!selectedConfig.value.name) {
      return;
    }
    const formData = new FormData();
    formData.append("fileName", `${selectedConfig.value.name}`);
    const res = await axios.post(API.REMOVE_CONF_NGINX, formData);
    if (res.success) {
      getListConf();
    }

    dataLogs.value.push(`[${new Date().toLocaleTimeString()}] ${res.message}`);
  } catch (error) {
    console.log(`Nginx: ${error}`);
  }
}

async function saveFileConf() {
  try {
    if (!selectedConfig.value.name || !contentFileConf.value) {
      return;
    }
    const formData = new FormData();
    formData.append("fileName", `${selectedConfig.value.name}`);
    formData.append("content", `${contentFileConf.value.replace(/\$/g, '\\$')}`);
    console.log(contentFileConf.value.replace(/\$/g, '\\$'))
    const res = await axios.post(API.SAVE_CONF_NGINX, formData);

    dataLogs.value.push(`[${new Date().toLocaleTimeString()}] ${res.message}`);
  } catch (error) {
    console.log(`Nginx: ${error}`);
  }
}

async function checkConf() {
  try {
    const res = await axios.get(`${API.CHECK_CONF_NGINX}`);
    dataLogs.value.push(`[${new Date().toLocaleTimeString()}] ${res.message}`);
  } catch (error) {
    console.log(`Nginx: ${error}`);
  }
}
</script>

<style scoped>
.btn-action {
  @apply flex items-center justify-center px-4 py-2 text-white rounded-lg transition-colors duration-200;
}

.btn-action:disabled {
  @apply opacity-50 cursor-not-allowed;
}
</style>

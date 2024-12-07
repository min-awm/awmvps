<template>
  <div class="p-3 m-2 bg-white border rounded-md">
    <div class="p-6" v-if="step == 0">
      <h2 class="mb-4 text-2xl font-semibold">MariaDB</h2>

      <div class="flex items-center space-x-4">
        <div class="w-4 h-4 bg-gray-500 rounded-full"></div>
        <span class="text-lg">Chưa cài đặt</span>
      </div>
      <button
        @click="installMariaDB"
        class="mt-4 bg-green-500 btn-action hover:bg-green-600"
        :disabled="isProcessingInstall"
      >
        <Download v-if="!isProcessingInstall" class="w-5 h-5 mr-2" />
        <Loader2 v-else class="w-5 h-5 mr-2 animate-spin" />
        Cài đặt
      </button>
    </div>

    <div class="p-6" v-if="step == 1">
      <h2 class="mb-4 text-2xl font-semibold">MariaDB</h2>

      <div class="flex items-center mb-3 space-x-4">
        <span class="text-lg">{{ version }}</span>
      </div>
      <div class="flex items-center mb-3 space-x-4">
        <div :class="`w-4 h-4 ${colorStatus(status)} rounded-full`"></div>
        <span class="text-lg">{{ textStatus(status) }}</span>
        <button
          :class="` btn-action cursor-pointer ${
            status === 'active\n' ? 'bg-red-500' : 'bg-green-500'
          }`"
          :disabled="isProcessingStopStart"
          @click="stopStartMariaDB(status)"
        >
          <Loader2
            class="w-5 h-5 mr-2 animate-spin"
            v-if="isProcessingStopStart"
          />
          {{ status === "active\n" ? "Dừng" : "Chạy" }}
        </button>
      </div>
      <div class="grid grid-cols-1 gap-4 mb-8 md:grid-cols-4"></div>

      <div class="">
        <div class="px-3 py-4 border-b border-gray-200 sm:px-6 bg-gray-50">
          <h1 class="text-xl font-semibold text-gray-900">
            Danh sách người dùng
          </h1>
        </div>
        <ul class="overflow-y-auto divide-y divide-gray-200 max-h-52">
          <li
            v-for="(user, i) in users"
            :key="`Mariadb-users-${i}`"
            class="px-4 py-4 transition-colors duration-150 sm:px-6 hover:bg-gray-50"
          >
            <div
              class="flex flex-wrap items-center justify-between sm:flex-nowrap"
            >
              <div class="flex items-center space-x-3">
                <h2 class="text-lg font-medium text-gray-900">
                  {{ user.User }}
                </h2>
              </div>
              <div class="flex items-center space-x-3">
                <h2 class="text-lg font-medium text-gray-900">
                  {{ user.Host }}
                </h2>
              </div>
              <div class="flex items-center mt-2 space-x-2 sm:mt-0">
                <button
                  @click="deleteUser(user.User, user.Host)"
                  class="px-3 py-1 text-white transition-colors duration-150 bg-red-500 rounded-md hover:bg-red-600 focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-offset-2"
                  :disabled="isProcessingDeleteUser"
                >
                  <Loader2
                    class="w-5 h-5 mr-2 animate-spin"
                    v-if="isProcessingDeleteUser"
                  />
                  Delete
                </button>
              </div>
            </div>
          </li>
        </ul>
        <div class="px-4 py-4 border-t border-gray-200 sm:px-6 bg-gray-50">
          <button
            @click="showCreateUserModal = true"
            class="w-full px-4 py-2 text-white transition-colors duration-150 bg-green-500 rounded-md sm:w-auto hover:bg-green-600 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2"
          >
            Thêm người dùng mới
          </button>
        </div>
      </div>
    </div>

    <div class="p-6 pt-0">
      <h2 class="mb-4 text-2xl font-semibold">Logs</h2>
      <div
        class="h-48 p-4 overflow-y-auto font-mono text-sm text-green-400 bg-gray-900 rounded-lg"
      >
        <p v-for="(log, i) in mysqlLogs" :key="`Mariadb-log-${i}`">{{ log }}</p>
      </div>
    </div>

    <div class="p-6 py-0" v-if="step == 1">
      <button
        @click="removeMariaDB()"
        class="bg-red-500 btn-action hover:bg-red-600"
        :disabled="isProcessingRemove"
      >
        <Trash2 v-if="!isProcessingRemove" class="w-5 h-5 mr-2" />
        <Loader2 v-else class="w-5 h-5 mr-2 animate-spin" />
        Gỡ cài đặt
      </button>
    </div>
  </div>

  <!-- Modal add user -->
  <div
    class="fixed inset-0 z-10 flex items-center justify-center p-4 bg-black bg-opacity-50"
    v-if="showCreateUserModal"
  >
    <div class="w-full max-w-md p-6 bg-white rounded-lg">
      <h1 class="mb-6 text-2xl font-medium text-gray-900">Thêm người dùng</h1>

      <form @submit.prevent="handleCreateUserSubmit" class="space-y-4">
        <div>
          <label
            for="username"
            class="block mb-1 text-sm font-medium text-gray-700"
          >
            Tên đăng nhập
          </label>
          <div
            class="flex items-center px-3 border border-gray-300 rounded-md shadow-sm"
          >
            <p>awmvps_</p>
            <input
              type="text"
              id="username"
              v-model="formUserData.username"
              class="w-full py-2 outline-none"
              placeholder="Nhập tên đăng nhập"
              required
            />
          </div>
        </div>

        <div>
          <label
            for="password"
            class="block mb-1 text-sm font-medium text-gray-700"
          >
            Mật khẩu
          </label>
          <input
            type="password"
            id="passwordUser"
            v-model="formUserData.password"
            class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm outline-none"
            placeholder="Nhập mật khẩu"
            required
          />
        </div>

        <div>
          <label
            for="passwordConfirm"
            class="block mb-1 text-sm font-medium text-gray-700"
          >
            Nhập lại mật khẩu
          </label>
          <input
            type="password"
            id="passwordUserConfirm"
            v-model="formUserData.passwordConfirm"
            class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm outline-none"
            placeholder="Nhập lại mật khẩu"
            required
          />
        </div>

        <div>
          <label
            for="passwordConfirm"
            class="block mb-1 text-sm font-medium text-gray-700"
          >
            IP
          </label>
          <input
            type="text"
            id="ip"
            v-model="formUserData.ip"
            class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm outline-none"
            placeholder="Nhập ip"
            required
          />
          <p class="mt-2 text-sm text-red-500">Note: % cho tất cả ip</p>
        </div>

        <div class="flex items-center justify-end mt-6 space-x-3">
          <button
            type="button"
            @click="showCreateUserModal = false"
            class="px-4 py-2 text-sm font-medium text-gray-700 border border-gray-300 rounded-md hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          >
            Hủy
          </button>
          <button
            type="submit"
            class="px-4 py-2 text-sm font-medium text-white bg-blue-600 border border-transparent rounded-md shadow-sm hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
            :disabled="isCreateUserSubmitting"
          >
            <span v-if="isCreateUserSubmitting" class="flex items-center">
              <Loader2Icon class="w-4 h-4 mr-2 -ml-1 animate-spin" />
              Đang xử lý...
            </span>
            <span v-else>Xác nhận</span>
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { Download, Loader2, Loader2Icon, Trash2 } from "lucide-vue-next";
import axios from "@/axios";
import API from "@/api";
import { useToast } from "vue-toastification";

const toast = useToast();
const step = ref(0);
const mysqlLogs = ref([]);
const version = ref("");
const status = ref("");
const isProcessingInstall = ref(false);
const isProcessingStopStart = ref(false);
const users = ref([]);
const formUserData = ref({
  username: "",
  password: "",
  passwordConfirm: "",
  ip: "%",
});
const isCreateUserSubmitting = ref(false);
const showCreateUserModal = ref(false);
const isProcessingDeleteUser = ref(false);
const isProcessingRemove = ref(false);

checkStatus();
getListUser();

async function checkStatus() {
  const installingMariaDB = localStorage.getItem("installingMariaDB");
  if (installingMariaDB == "true") {
    isProcessingInstall.value = true;
  }

  const removingMariaDB = localStorage.getItem("removingMariaDB");

  if (removingMariaDB == "true") {
    isProcessingRemove.value = true;
  }

  try {
    const res = await axios.get(`${API.STATUS_MARIADB}`);
    if (res.success) {
      version.value = res.version;
      status.value = res.status;
      step.value = 1;
      localStorage.setItem("installingMariaDB", false);
    } else {
      switch (res.message) {
        case "notInstall":
          step.value = 0;
          localStorage.setItem("removingMariaDB", false);
          break;

        case "versionErr":
          step.value = 1;
          mysqlLogs.value.push(
            `[${new Date().toLocaleTimeString()}] Lỗi: Không tìm thấy version MariaDB`
          );
          break;

        default:
          break;
      }
    }
  } catch (error) {
    console.log(`Mariadb: ${error}`);
  }
}

async function installMariaDB() {
  isProcessingInstall.value = true;
  mysqlLogs.value.push(
    `[${new Date().toLocaleTimeString()}] Quá trình cài đặt có thể mất vài phút.`
  );
  localStorage.setItem("installingMariaDB", true);
  try {
    const res = await axios.get(`${API.INSTALL_MARIADB}`);
    if (res.success) {
      checkStatus();
    }
    mysqlLogs.value.push(`[${new Date().toLocaleTimeString()}] ${res.message}`);
  } catch (error) {
    console.log(`Mariadb: ${error}`);
  } finally {
    isProcessingInstall.value = false;
    localStorage.setItem("installingMariaDB", false);
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

async function stopStartMariaDB(status) {
  isProcessingStopStart.value = true;
  let url = API.START_MARIADB;
  if (status == "active\n") {
    url = API.STOP_MARIADB;
  }

  try {
    const res = await axios.get(`${url}`);
    if (res.success) {
      checkStatus();
    }

    mysqlLogs.value.push(`[${new Date().toLocaleTimeString()}] ${res.message}`);
  } catch (error) {
    console.log(`Mariadb: ${error}`);
  } finally {
    isProcessingStopStart.value = false;
  }
}

async function getListUser() {
  try {
    const res = await axios.get(`${API.LIST_USER_MARIADB}`);
    if (res.success) {
      const input = res.message;
      const lines = input.trim().split("\n");
      const headers = lines[0].split("\t");

      const result = lines.slice(1).map((line) => {
        const values = line.split("\t");
        return headers.reduce((obj, header, index) => {
          obj[header] = values[index];
          return obj;
        }, {});
      });

      users.value = result;
    }
  } catch (error) {
    console.log(`Mariadb: ${error}`);
  }
}

async function handleCreateUserSubmit() {
  if (formUserData.value.password !== formUserData.value.passwordConfirm) {
    return toast.warning("Mật khẩu không giống nhau");
  }
  isCreateUserSubmitting.value = true;

  try {
    const formData = new FormData();
    formData.append("newUser", `awmvps_${formUserData.value.username}`);
    formData.append("password", formUserData.value.password);
    formData.append("ip", formUserData.value.ip);

    const res = await axios.post(API.CREATE_USER_MARIADB, formData);
    if (res.success) {
      getListUser();
      toast.success("Đã thêm tài khoản thành công");
      mysqlLogs.value.push(
        `[${new Date().toLocaleTimeString()}] Đã thêm tài khoản ${
          formUserData.value.username
        }@${formUserData.value.ip}`
      );
    } else {
      mysqlLogs.value.push(
        `[${new Date().toLocaleTimeString()}] ${res.message}`
      );
    }
  } catch (error) {
    console.log(`Mariadb: ${error}`);
  } finally {
    showCreateUserModal.value = false;
    isCreateUserSubmitting.value = false;
  }
}

async function deleteUser(user, host) {
  isProcessingDeleteUser.value = true;
  try {
    const formData = new FormData();
    formData.append("newUser", user);
    formData.append("ip", host);

    const res = await axios.post(API.DELETE_USER_MARIADB, formData);
    if (res.success) {
      getListUser();
      toast.success("Đã xoá tài khoản thành công");
      mysqlLogs.value.push(
        `[${new Date().toLocaleTimeString()}] Đã xoá tài khoản ${user}@${host}`
      );
    } else {
      mysqlLogs.value.push(
        `[${new Date().toLocaleTimeString()}] ${res.message}`
      );
    }
  } catch (error) {
    console.log(`Mariadb: ${error}`);
  } finally {
    isProcessingDeleteUser.value = false;
  }
}

async function removeMariaDB() {
  isProcessingRemove.value = true;
  mysqlLogs.value.push(
    `[${new Date().toLocaleTimeString()}] Quá trình gỡ cài đặt có thể mất vài phút.`
  );
  localStorage.setItem("removingMariaDB", true);

  try {
    const res = await axios.get(`${API.REMOVE_MARIADB}`);
    if (res.success) {
      checkStatus();
    }

    mysqlLogs.value.push(`[${new Date().toLocaleTimeString()}] ${res.message}`);
  } catch (error) {
    console.log(`Mariadb: ${error}`);
  } finally {
    isProcessingRemove.value = false;
    localStorage.setItem("removingMariaDB", false);
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

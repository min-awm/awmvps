<template>
  <div class="m-2 bg-white border rounded-md p-9">
    <!-- Certificate Auto-renewal Section -->
    <section class="mb-8">
      <h2 class="flex justify-between mb-4 text-xl font-bold text-gray-900">
        Chứng chỉ Ssl

        <OpenTerminalButton />
      </h2>
      <p class="mb-4 text-gray-600">
        Lưu ý: Chứng chỉ Let's Encrypt có hạn 90 ngày.
      </p>

      <div class="flex flex-col gap-3 mb-4">
        <p class="text-gray-900">Cài đặt Certbot:</p>
        <Code
          command="sudo apt install certbot python3-certbot-nginx"
          v-if="packageManager == 'not' || packageManager == 'apt'"
        />
        <Code
          command="sudo yum install certbot python3-certbot-nginx"
          v-if="packageManager == 'not' || packageManager == 'yum'"
        />
      </div>

      <div class="mb-4">
        <p class="mb-2 text-gray-900">
          Tạo chứng chỉ SSL cho tên miền(example.com):
        </p>
        <div class="flex flex-col gap-4">
          <Code command="sudo certbot --nginx -d example.com" />
        </div>
      </div>
    </section>

    <!-- HTTPS Verification Section -->
    <!-- <section>
      <h2 class="mb-4 text-xl font-bold text-gray-900">7. Kiểm tra HTTPS</h2>
      <ul class="space-y-2 text-gray-600 list-disc list-inside">
        <li>
          Truy cập website bằng HTTPS, ví dụ:
          <code class="bg-gray-100 px-2 py-0.5 rounded text-gray-800">
            https://example.com
          </code>
        </li>
        <li>
          Sử dụng công cụ kiểm tra SSL như
          <a
            href="https://www.ssllabs.com"
            class="text-blue-600 hover:text-blue-800 hover:underline"
          >
            SSL Labs
          </a>
          để xác minh.
        </li>
      </ul>
    </section> -->
  </div>
</template>

<script setup>
import { ref } from "vue";
import OpenTerminalButton from "@/components/common/OpenTerminalButton.vue";
import Code from "@/components/common/Code.vue";
import axios from "@/axios";
import API from "@/api";

const packageManager = ref("not");
getPackageManager()

async function getPackageManager() {
  try {
    const res = await axios.get(`${API.PACKAGE_MANAGER}`);
    if (res.success) {
      packageManager.value = res.data
    }

  } catch (error) {
    console.log(`Ssl: ${error}`);
  }
}
</script>

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

      <div class="mb-4">
        <p class="mb-2 text-gray-900">Cài đặt Certbot:</p>
        <Code command="sudo yum install certbot python3-certbot-nginx" />
      </div>

      <div class="mb-4">
        <p class="mb-2 text-gray-900">
          Tạo chứng chỉ SSL cho tên miền(example.com):
        </p>
        <Code command="sudo certbot --nginx -d example.com" />
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
import { Copy, Check } from "lucide-vue-next";
import OpenTerminalButton from "@/components/common/OpenTerminalButton.vue";
import Code from "@/components/common/Code.vue";

const copyStatus = ref("");

const copyToClipboard = async (text) => {
  try {
    await navigator.clipboard.writeText(text);
    copyStatus.value = text.includes("status") ? "status" : "enable";
    setTimeout(() => {
      copyStatus.value = "";
    }, 2000);
  } catch (err) {
    console.error("Failed to copy text: ", err);
  }
};
</script>

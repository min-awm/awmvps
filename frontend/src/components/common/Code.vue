<template>
  <div class="relative overflow-hidden border rounded-lg bg-gray-50">
    <div class="absolute right-2 top-2">
      <button
        @click="copyToClipboard(command)"
        class="p-1.5 hover:bg-gray-200 rounded-md transition-colors"
        :class="{ 'text-green-500': copyStatus === 'done' }"
      >
        <Check v-if="copyStatus === 'done'" class="w-4 h-4" />
        <ClipboardPaste v-else class="w-4 h-4 text-gray-500" />
      </button>
    </div>
    <div class="px-4 py-2 text-sm text-gray-600 bg-gray-100 border-b">bash</div>
    <pre class="p-4 overflow-x-auto text-sm"><code>{{ command }}</code></pre>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { ClipboardPaste, Check } from "lucide-vue-next";
import { useTerminalStore } from "@/store/terminal";

const terminalStore = useTerminalStore();
const props = defineProps(["command"]);
const copyStatus = ref("");

async function copyToClipboard(text) {
  try {
    terminalStore.setContentCommand(text)
    await navigator.clipboard.writeText(text);
    copyStatus.value = "done";
    setTimeout(() => {
      copyStatus.value = "";
    }, 2000);
  } catch (err) {
    console.error("Failed to copy text: ", err);
  }
}
</script>

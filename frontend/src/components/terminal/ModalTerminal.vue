<template>
  <!-- bg-[#1e1e1e] -->
  <div
    class="w-full max-w-4xl bg-[#fff] rounded-md shadow-xl overflow-hidden z-[9999] right-0 bottom-0 fixed"
    :style="{ transform: `translate(${position.x}px, ${position.y}px)` }"
  >
    <!-- Terminal Header -->
    <div
      class="bg-[#3c3c3c] px-4 py-2 flex items-center justify-between cursor-move"
      @mousedown="startDrag"
      @touchstart.passive="startDrag"
    >
      <div class="flex items-center space-x-2">
        <span class="text-sm text-gray-300">FE - Google Chrome</span>
      </div>
      <div class="flex items-center space-x-2">
        <button class="p-1 hover:bg-[#4a4a4a] rounded" @click="closeTerminal">
          <X class="w-4 h-4 text-gray-400" />
        </button>
      </div>
    </div>

    <!-- Terminal Content -->
    <Xterm />
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from "vue";
import { X } from "lucide-vue-next";
import Xterm from "@/components/terminal/Xterm.vue";
import { useTerminalStore } from "@/store/terminal";

const terminalStore = useTerminalStore();

const position = ref({ x: 0, y: 0 });
let isDragging = false;
let startPosition = { x: 0, y: 0 };

function startDrag(event) {
  isDragging = true;
  startPosition = {
    x: event.clientX || event.touches[0].clientX,
    y: event.clientY || event.touches[0].clientY,
  };
}

function drag(event) {
  if (!isDragging) return;
  const currentPosition = {
    x: event.clientX || event.touches[0].clientX,
    y: event.clientY || event.touches[0].clientY,
  };
  position.value = {
    x: position.value.x + (currentPosition.x - startPosition.x),
    y: position.value.y + (currentPosition.y - startPosition.y),
  };
  startPosition = currentPosition;
}

function stopDrag() {
  isDragging = false;
}

function closeTerminal() {
  terminalStore.setShowTerminal(false);
}

onMounted(() => {
  document.addEventListener("mousemove", drag);
  document.addEventListener("mouseup", stopDrag);
  document.addEventListener("touchmove", drag);
  document.addEventListener("touchend", stopDrag);
});

onUnmounted(() => {
  document.removeEventListener("mousemove", drag);
  document.removeEventListener("mouseup", stopDrag);
  document.removeEventListener("touchmove", drag);
  document.removeEventListener("touchend", stopDrag);
});
</script>

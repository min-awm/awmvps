<template>
    <div class="flex items-center justify-center min-h-screen p-4 bg-gray-100">
      <div 
        v-drag
        class="w-full max-w-4xl bg-[#2f2f2f] rounded-md shadow-xl overflow-hidden cursor-move"
        :style="{ transform: `translate(${position.x}px, ${position.y}px)` }"
      >
        <!-- Terminal Header -->
        <div 
          class="bg-[#3c3c3c] px-4 py-2 flex items-center justify-between"
          @mousedown="startDrag"
          @touchstart="startDrag"
        >
          <div class="flex items-center space-x-2">
            <span class="text-sm text-gray-300">FE - Google Chrome</span>
          </div>
          <div class="flex items-center space-x-2">
            <button class="p-1 hover:bg-[#4a4a4a] rounded">
              <Minus class="w-4 h-4 text-gray-400" />
            </button>
            <button class="p-1 hover:bg-[#4a4a4a] rounded">
              <Square class="w-4 h-4 text-gray-400" />
            </button>
            <button class="p-1 hover:bg-[#4a4a4a] rounded">
              <X class="w-4 h-4 text-gray-400" />
            </button>
          </div>
        </div>
  
        <!-- URL Bar -->
        <div class="bg-[#2b2b2b] px-4 py-2 flex items-center space-x-2">
          <div class="flex items-center flex-1">
            <InfoIcon class="w-4 h-4 mr-2 text-gray-500" />
            <div class="flex items-center bg-[#242424] rounded px-2 py-1 flex-1">
              <span class="text-sm text-gray-400">localhost:</span>
              <span class="text-sm text-gray-300">5173/#/mini-terminal</span>
            </div>
            <RefreshCw class="w-4 h-4 ml-2 text-gray-500 cursor-pointer hover:text-gray-300" />
          </div>
        </div>
  
        <!-- Terminal Content -->
        <div class="h-[400px] p-4 font-mono text-sm overflow-auto">
          <div class="flex items-start">
            <span class="text-green-500">root@fedora</span>
            <span class="text-gray-400">:</span>
            <span class="text-blue-400">/home/minh/my-code/awmvps</span>
            <span class="text-gray-400">#</span>
            <span class="ml-1 animate-blink">|</span>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref, onMounted, onUnmounted } from 'vue'
  import { Minus, Square, X, RefreshCw, InfoIcon } from 'lucide-vue-next'
  
  const position = ref({ x: 0, y: 0 })
  let isDragging = false
  let startPosition = { x: 0, y: 0 }
  
  const startDrag = (event) => {
    isDragging = true
    startPosition = {
      x: event.clientX || event.touches[0].clientX,
      y: event.clientY || event.touches[0].clientY
    }
  }
  
  const drag = (event) => {
    if (!isDragging) return
    const currentPosition = {
      x: event.clientX || event.touches[0].clientX,
      y: event.clientY || event.touches[0].clientY
    }
    position.value = {
      x: position.value.x + (currentPosition.x - startPosition.x),
      y: position.value.y + (currentPosition.y - startPosition.y)
    }
    startPosition = currentPosition
  }
  
  const stopDrag = () => {
    isDragging = false
  }
  
  onMounted(() => {
    document.addEventListener('mousemove', drag)
    document.addEventListener('mouseup', stopDrag)
    document.addEventListener('touchmove', drag)
    document.addEventListener('touchend', stopDrag)
  })
  
  onUnmounted(() => {
    document.removeEventListener('mousemove', drag)
    document.removeEventListener('mouseup', stopDrag)
    document.removeEventListener('touchmove', drag)
    document.removeEventListener('touchend', stopDrag)
  })
  </script>
  
  <style scoped>
  @keyframes blink {
    0%, 100% { opacity: 1; }
    50% { opacity: 0; }
  }
  
  .animate-blink {
    animation: blink 1s step-end infinite;
    color: #fff;
  }
  
  /* Mobile Responsiveness */
  @media (max-width: 640px) {
    .min-h-screen {
      padding: 0.5rem;
    }
    
    .text-sm {
      font-size: 0.75rem;
    }
    
  }
  
  /* Custom Scrollbar */
  .overflow-auto::-webkit-scrollbar {
    width: 8px;
    height: 8px;
  }
  
  .overflow-auto::-webkit-scrollbar-track {
    background: #2b2b2b;
  }
  
  .overflow-auto::-webkit-scrollbar-thumb {
    background: #4a4a4a;
    border-radius: 4px;
  }
  
  .overflow-auto::-webkit-scrollbar-thumb:hover {
    background: #5a5a5a;
  }
  
  /* Disable text selection during drag */
  .cursor-move {
    user-select: none;
  }
  </style>
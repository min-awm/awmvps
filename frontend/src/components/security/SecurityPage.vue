<template>
    <div class="w-full m-10 bg-white rounded-lg shadow">
      <!-- Header -->
      <button 
        @click="isExpanded = !isExpanded"
        class="flex items-center justify-between w-full p-4 hover:bg-gray-50"
        :aria-expanded="isExpanded"
      >
        <div class="flex items-center space-x-2">
          <Shield class="w-6 h-6 text-gray-700" />
          <h2 class="text-lg font-medium text-gray-900">Security</h2>
        </div>
        <ChevronUp 
          class="w-5 h-5 text-gray-500 transition-transform duration-200"
          :class="{ 'rotate-180': !isExpanded }"
        />
      </button>
  
      <!-- Grid Content -->
      <div 
        v-show="isExpanded"
        class="p-4"
      >
        <div class="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-3">
          <!-- SSH Access -->
          <button class="flex items-center p-3 space-x-3 transition-colors rounded-lg hover:bg-gray-50">
            <div class="flex items-center justify-center flex-shrink-0 w-10 h-10 rounded-lg bg-orange-50">
              <Terminal class="w-5 h-5 text-orange-600" />
            </div>
            <span class="text-sm font-medium text-gray-900">SSH Access</span>
          </button>
  
          <!-- IP Blocker -->
          <button class="flex items-center p-3 space-x-3 transition-colors rounded-lg hover:bg-gray-50">
            <div class="flex items-center justify-center flex-shrink-0 w-10 h-10 rounded-lg bg-blue-50">
              <Shield class="w-5 h-5 text-blue-600" />
            </div>
            <span class="text-sm font-medium text-gray-900">IP Blocker</span>
          </button>
  
          <!-- SSL/TLS -->
          <button class="flex items-center p-3 space-x-3 transition-colors rounded-lg hover:bg-gray-50">
            <div class="flex items-center justify-center flex-shrink-0 w-10 h-10 rounded-lg bg-green-50">
              <Lock class="w-5 h-5 text-green-600" />
            </div>
            <span class="text-sm font-medium text-gray-900">SSL/TLS</span>
          </button>
  
          <!-- Manage API Tokens -->
          <button class="flex items-center p-3 space-x-3 transition-colors rounded-lg hover:bg-gray-50">
            <div class="flex items-center justify-center flex-shrink-0 w-10 h-10 rounded-lg bg-purple-50">
              <Key class="w-5 h-5 text-purple-600" />
            </div>
            <span class="text-sm font-medium text-gray-900"> Brute Force Minitor</span>
          </button>
  
          <!-- Hotlink Protection -->
          <button class="flex items-center p-3 space-x-3 transition-colors rounded-lg hover:bg-gray-50">
            <div class="flex items-center justify-center flex-shrink-0 w-10 h-10 rounded-lg bg-red-50">
              <Link class="w-5 h-5 text-red-600" />
            </div>
            <span class="text-sm font-medium text-gray-900">Hotlink Protection</span>
          </button>
  
          <!-- Leech Protection -->
          <button class="flex items-center p-3 space-x-3 transition-colors rounded-lg hover:bg-gray-50">
            <div class="flex items-center justify-center flex-shrink-0 w-10 h-10 rounded-lg bg-yellow-50">
              <Shield class="w-5 h-5 text-yellow-600" />
            </div>
            <span class="text-sm font-medium text-gray-900">Leech Protection</span>
          </button>
  
          <!-- SSL/TLS Status -->
          <button class="flex items-center p-3 space-x-3 transition-colors rounded-lg hover:bg-gray-50">
            <div class="flex items-center justify-center flex-shrink-0 w-10 h-10 rounded-lg bg-cyan-50">
              <FileCheck class="w-5 h-5 text-cyan-600" />
            </div>
            <span class="text-sm font-medium text-gray-900">SSL/TLS Status</span>
          </button>
  
          <!-- Two-Factor Authentication -->
          <button class="flex items-center p-3 space-x-3 transition-colors rounded-lg hover:bg-gray-50">
            <div class="flex items-center justify-center flex-shrink-0 w-10 h-10 rounded-lg bg-indigo-50">
              <Smartphone class="w-5 h-5 text-indigo-600" />
            </div>
            <span class="text-sm font-medium text-gray-900">Check IP</span>
          </button>
  
          <!-- Imunify360 -->
          <button class="flex items-center p-3 space-x-3 transition-colors rounded-lg hover:bg-gray-50">
            <div class="flex items-center justify-center flex-shrink-0 w-10 h-10 rounded-lg bg-rose-50">
              <ShieldCheck class="w-5 h-5 text-rose-600" />
            </div>
            <span class="text-sm font-medium text-gray-900">DDOS</span>
          </button>
        </div>
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue'
  import {
    Shield,
    ChevronUp,
    Terminal,
    Lock,
    Key,
    Link,
    FileCheck,
    Smartphone,
    ShieldCheck
  } from 'lucide-vue-next'
  
  const isExpanded = ref(true)
  </script>
  
  <style scoped>
  /* Add any additional component-specific styles here if needed */
  </style>
<template>
  <router-view />

  <ModalTerminal v-if="showTerminal" />
</template>

<script setup>
import ModalTerminal from "@/components/terminal/ModalTerminal.vue";
import { useTerminalStore } from '@/store/terminal'
import { storeToRefs } from 'pinia'

const terminalStore = useTerminalStore()
const { showTerminal } = storeToRefs(terminalStore)
</script>
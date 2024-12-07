import { ref } from "vue";
import { defineStore } from "pinia";

export const useTerminalStore = defineStore("terminal", () => {
  const showTerminal = ref(false);
  const terminalWriteFn = ref(null);

  function setShowTerminal(value) {
    showTerminal.value = value;
  }

  function setContentCommand(value) {
    showTerminal.value = true;
    setTimeout(() => {
      if (!terminalWriteFn.value) return;
      terminalWriteFn.value(value);
    }, 500);
  }

  function setTerminalWriteFn(fn) {
    terminalWriteFn.value = fn;
  }

  return {
    showTerminal,
    setShowTerminal,
    setContentCommand,
    setTerminalWriteFn,
  };
});

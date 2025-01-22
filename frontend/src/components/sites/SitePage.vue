<template>
  <div class="m-2 bg-white border rounded-md p-9">
    <!-- Tabs and Create New Button -->
    <div
      class="flex flex-col mb-4 sm:flex-row sm:items-center sm:justify-between"
    >
      <div class="mb-4 border-b sm:mb-0 sm:border-b-0">
        <div class="flex space-x-8">
          <button
            class="px-0 pb-2 text-sm font-medium border-b-2 border-[#0070f3] text-[#0070f3]"
          >
            Website
          </button>
        </div>
      </div>
      <button
        @click="createNewWorkspace"
        class="bg-[#0070f3] text-white hover:bg-[#0060df] px-4 py-2 rounded-md text-sm font-medium"
      >
        Thêm mới
      </button>
    </div>

    <!-- Workspace Cards -->
    <div class="space-y-2">
      <div
        v-for="(workspace, i) in workspaces"
        :key="`SitePage-workspaces-${i}`"
        class="bg-white rounded-lg border border-[#eaeaea] p-4 flex items-center shadow-sm"
      >
        <!-- Workspace Info -->
        <div class="flex items-center flex-1 min-w-0">
          <!-- Logo -->
          <div class="flex items-center justify-center pl-4 pr-6">
            {{ i + 1 }}
          </div>

          <!-- Text Content -->
          <div class="min-w-0">
            <h3 class="font-medium text-sm text-[#000]">
              {{ workspace.name }}
            </h3>
            <div class="flex items-center gap-2 text-sm text-[#666]">
              <span class="truncate">{{ workspace.id }}</span>
              <span
                v-if="workspace.status"
                class="inline-flex items-center rounded-full px-2 py-0.5 text-xs bg-[#eaeaea] text-[#666]"
              >
                {{ workspace.status }}
              </span>
            </div>
          </div>
        </div>

        <!-- Dropdown Menu -->
        <div class="relative">
          <button
            @click="toggleDropdown(i)"
            class="ml-2 p-2 hover:bg-[#f6f6f6] rounded-full"
          >
            <MoreHorizontalIcon class="h-4 w-4 text-[#666]" />
            <span class="sr-only">Open menu</span>
          </button>

          <!-- Dropdown Content -->
          <div
            v-if="activeDropdown == i"
            class="absolute right-0 mt-2 w-48 py-1 bg-white rounded-md shadow-lg border border-[#eaeaea] z-10"
          >
            <button
              v-for="action in dropdownActions"
              :key="action"
              class="w-full px-4 py-2 text-sm text-left hover:bg-[#f6f6f6] text-[#000]"
              @click="handleAction(action, workspace)"
            >
              {{ action }}
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { MoreHorizontalIcon } from "lucide-vue-next";

// Workspaces data
const workspaces = ref([
  {
    id: "test-3345368",
    name: "test",
    logo: "https://hebbkx1anhila5yf.public.blob.vercel-storage.com/image-NCMMB7FLh6LXJgdGXct2dFND3oHBir.png",
    status: "Archived",
  },
  {
    id: "test-3345368",
    name: "test",
    logo: "https://hebbkx1anhila5yf.public.blob.vercel-storage.com/image-NCMMB7FLh6LXJgdGXct2dFND3oHBir.png",
    status: "Archived",
  },
  {
    id: "test-3345368",
    name: "test",
    logo: "https://hebbkx1anhila5yf.public.blob.vercel-storage.com/image-NCMMB7FLh6LXJgdGXct2dFND3oHBir.png",
    status: "Archived",
  },
]);

// Dropdown functionality
const activeDropdown = ref(null);
const dropdownActions = ["View details", "Unarchive", "Delete"];

const toggleDropdown = (i) => {
  activeDropdown.value = activeDropdown.value === i ? null : i;
};

const closeDropdown = () => {
  activeDropdown.value = null;
};

const handleAction = (action, workspace) => {
  console.log(`${action} clicked for workspace:`, workspace.name);
  closeDropdown();
};

// Create New Workspace functionality
const createNewWorkspace = () => {
  console.log("Create New Workspace clicked");
  // Add your logic here to create a new workspace
};
</script>

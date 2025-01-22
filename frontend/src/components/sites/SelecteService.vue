<template>
  <nav class="w-[100px]">
    <ul class="space-y-2">
      <li v-for="(item, index) in menuItems" :key="index">
        <div class="space-y-1">
          <button
            @click="toggleItem(index)"
            class="flex items-center justify-between w-full p-4 transition-colors duration-200 rounded-lg hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200"
            :class="{
              'bg-blue-50 text-blue-900': selectedIndex === index,
              'text-gray-700': selectedIndex !== index,
            }"
            :aria-expanded="expandedItems[index]"
            :aria-controls="`submenu-${index}`"
          >
            <span class="text-sm font-medium">{{ item.name }}</span>
            <div class="flex items-center space-x-2">
              <span v-if="item.children?.length" class="text-xs text-gray-500">
                {{ item.children.length }}
              </span>
              <ChevronRightIcon
                class="w-5 h-5 transition-transform duration-200"
                :class="{
                  'text-blue-500 rotate-90': expandedItems[index],
                  'text-gray-400': !expandedItems[index],
                }"
              />
            </div>
          </button>

          <!-- Submenu -->
          <transition
            enter-active-class="transition duration-100 ease-out"
            enter-from-class="transform scale-95 opacity-0"
            enter-to-class="transform scale-100 opacity-100"
            leave-active-class="transition duration-75 ease-in"
            leave-from-class="transform scale-100 opacity-100"
            leave-to-class="transform scale-95 opacity-0"
          >
            <ul
              v-if="expandedItems[index] && item.children?.length"
              :id="`submenu-${index}`"
              class="pl-4"
            >
              <li
                v-for="(child, childIndex) in item.children"
                :key="childIndex"
              >
                <button
                  @click="selectChild(index, childIndex)"
                  class="flex items-center justify-between w-full p-3 transition-colors duration-200 rounded-lg hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200"
                  :class="{
                    'bg-blue-50/50 text-blue-900':
                      selectedChildIndex === childIndex &&
                      selectedIndex === index,
                    'text-gray-600': !(
                      selectedChildIndex === childIndex &&
                      selectedIndex === index
                    ),
                  }"
                >
                  <span class="text-sm">{{ child.name }}</span>
                  <Plus
                    class="w-4 h-4"
                    :class="{
                      'text-blue-500':
                        selectedChildIndex === childIndex &&
                        selectedIndex === index,
                      'text-gray-400': !(
                        selectedChildIndex === childIndex &&
                        selectedIndex === index
                      ),
                    }"
                  />
                </button>
              </li>
            </ul>
          </transition>
        </div>
      </li>
    </ul>
  </nav>
</template>

<script setup>
import { ref } from "vue";
import { ChevronRightIcon, Plus } from "lucide-vue-next";

const emit = defineEmits(['selected'])

const menuItems = [
  {
    name: "Backend",
    children: [{ name: "Nodejs" }, { name: "Php" }, { name: "Python" }],
  },
  {
    name: "Databases",
    children: [{ name: "Git" }, { name: "Docker" }],
  },
];

const selectedIndex = ref(0);
const selectedChildIndex = ref(null);
const expandedItems = ref({});

const toggleItem = (index) => {
  expandedItems.value[index] = !expandedItems.value[index];
  selectedIndex.value = index;
};

const selectChild = (parentIndex, childIndex) => {
  selectedIndex.value = parentIndex;
  selectedChildIndex.value = childIndex;
  emit('selected', menuItems[parentIndex].children[childIndex])
};
</script>

<style scoped>
@media (max-width: 640px) {
  button {
    padding: 0.75rem;
  }
}

@media (min-width: 768px) {
  nav {
    min-width: 320px;
  }
}
</style>

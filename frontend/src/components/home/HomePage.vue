<template>
  <div class="min-h-screen bg-gray-50">
    <!-- Header with close button -->
    <header class="flex items-center justify-between p-4 bg-white border-b">
      <h1 class="text-lg font-semibold">Installing — NAME APP 2024 — 17.4.4</h1>
      <button class="p-2 rounded hover:bg-gray-100">
        <X class="w-5 h-5" />
      </button>
    </header>

    <!-- Tabs -->
    <div class="bg-white border-b">
      <nav class="flex overflow-x-auto">
        <button 
          v-for="tab in tabs" 
          :key="tab"
          :class="[
            'px-6 py-3 whitespace-nowrap',
            activeTab === tab ? 'border-b-2 border-blue-600 text-blue-600' : 'text-gray-600'
          ]"
          @click="activeTab = tab"
        >
          {{ tab }}
        </button>
      </nav>
    </div>

    <!-- Main content -->
    <div class="flex flex-col md:flex-row h-[calc(100vh-8rem)]">
      <!-- Left panel -->
      <div class="w-full p-4 overflow-auto md:w-3/5">
        <!-- Info banner -->
        <div class="flex items-center justify-between p-4 mb-4 rounded bg-blue-50">
          <div class="flex items-center">
            <InfoIcon class="w-5 h-5 mr-2 text-blue-600" />
            <span>Need help choosing what to install? <a href="#" class="text-blue-600">More info</a></span>
          </div>
          <button class="text-gray-500 hover:text-gray-700">
            <X class="w-5 h-5" />
          </button>
        </div>

        <!-- Workload sections -->
        <div class="space-y-6">
          <section v-for="(section, index) in workloads" :key="index">
            <h2 class="mb-3 text-lg font-semibold">{{ section.title }} ({{ section.items.length }})</h2>
            <div class="grid grid-cols-1 gap-4 lg:grid-cols-2">
              <div 
                v-for="item in section.items" 
                :key="item.title"
                class="p-4 transition-colors bg-white border rounded-lg hover:border-blue-400"
              >
                <div class="flex items-start space-x-4">
                  <input 
                    type="checkbox" 
                    :checked="item.selected"
                    class="w-5 h-5 mt-1 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
                  />
                  <div>
                    <div class="flex items-center space-x-2">
                      <component :is="item.icon" class="w-8 h-8 text-blue-600" />
                      <h3 class="font-medium">{{ item.title }}</h3>
                    </div>
                    <p class="mt-1 text-sm text-gray-600">{{ item.description }}</p>
                  </div>
                </div>
              </div>
            </div>
          </section>
        </div>
      </div>

      <!-- Right panel -->
      <div class="w-full overflow-auto bg-white border-l md:w-2/5">
        <div class="p-4">
          <h2 class="mb-4 text-lg font-semibold">Installation details</h2>
          <!-- Tree view -->
          <div class="space-y-2">
            <div v-for="(category, index) in installationDetails" :key="index">
              <div class="flex items-center space-x-2">
                <ChevronDown class="w-4 h-4" />
                <span class="font-medium">{{ category.title }}</span>
              </div>
              <div class="mt-2 ml-6 space-y-2">
                <div v-for="(item, itemIndex) in category.items" :key="itemIndex" class="flex items-center space-x-2">
                  <input 
                    type="checkbox" 
                    :checked="item.selected"
                    class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
                  />
                  <span class="text-sm">{{ item.title }}</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Footer -->
    <footer class="p-4 bg-white border-t">
      <div class="flex flex-col items-center justify-between space-y-4 md:flex-row md:space-y-0">
        <div class="text-sm text-gray-600">
          <span>Location: </span>
          <span class="font-mono">C:\Program Files\Microsoft Visual Studio\2022\Community</span>
          <button class="ml-2 text-blue-600">Change...</button>
        </div>
        <div class="flex space-x-4">
          <span class="text-sm text-gray-600">Total space required: 8.97 GB</span>
          <button class="px-4 py-2 text-white bg-blue-600 rounded hover:bg-blue-700">
            Install
          </button>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { 
  X, 
  Info as InfoIcon, 
  ChevronDown,
  Globe,
  Cloud,
  Code,
  Database,
  Monitor
} from 'lucide-vue-next'

const activeTab = ref('Workloads')

const tabs = [
  'Workloads',
  'Individual components',
  'Language packs',
  'Installation locations'
]

const workloads = [
  {
    title: 'Web & Cloud',
    items: [
      {
        icon: Globe,
        title: 'ASP.NET and web development',
        description: 'Build web applications using ASP.NET Core, ASP.NET, HTML/JavaScript, and Containers including Docker support',
        selected: true
      },
      {
        icon: Cloud,
        title: 'Azure development',
        description: 'Azure SDKs, tools, and projects for developing cloud apps and creating resources using .NET and .NET Framework',
        selected: false
      },
      {
        icon: Code,
        title: 'Python development',
        description: 'Editing, debugging, interactive development and source control for Python',
        selected: false
      },
      {
        icon: Database,
        title: 'Node.js development',
        description: 'Build scalable network applications using Node.js, an asynchronous event-driven JavaScript runtime',
        selected: false
      }
    ]
  },
  {
    title: 'Desktop & Mobile',
    items: [
      {
        icon: Monitor,
        title: '.NET Multi-platform App UI development',
        description: 'Build Android, iOS, Windows, and Mac apps from a single codebase using C# with .NET MAUI',
        selected: false
      },
      {
        icon: Monitor,
        title: '.NET desktop development',
        description: 'Build WPF, Windows Forms, and console applications using C#, Visual Basic, and F# with .NET and .NET Framework',
        selected: true
      }
    ]
  }
]

const installationDetails = [
  {
    title: '.NET desktop development',
    items: [
      { title: '.NET desktop development tools', selected: true },
      { title: '.NET Framework 4.7.2 development tools', selected: true },
      { title: 'C# and Visual Basic', selected: true },
      { title: 'Development tools for .NET', selected: true },
      { title: '.NET Framework 4.8 development tools', selected: true },
      { title: 'Entity Framework 6 tools', selected: true },
      { title: '.NET profiling tools', selected: true },
      { title: 'IntelliCode', selected: true },
      { title: 'Just-In-Time debugger', selected: true }
    ]
  }
]
</script>
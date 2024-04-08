<template>
  <div>
    <fieldset class="mt-4">
      <legend class="sr-only">Notification method</legend>
      <div class="space-y-4 sm:flex sm:items-center sm:space-x-10 sm:space-y-0">
        <div v-for="notificationMethod in notificationMethods" :key="notificationMethod.id" class="flex items-center">
          <input :id="notificationMethod.id" name="notification-method" type="radio" :checked="notificationMethod.id === 'email'" class="h-4 w-4 border-gray-300 text-indigo-600 focus:ring-indigo-600" />
          <label :for="notificationMethod.id" class="ml-3 block text-sm font-medium leading-6 text-gray-900">{{ notificationMethod.title }}</label>
        </div>
      </div>
    </fieldset>
  </div>
</template>

<script setup lang="ts">
const notificationMethods = [
  { id: 'purchasing', title: 'Purchasing', checked: true },
  { id: 'selling', title: 'Selling', checked: false },
]
</script>

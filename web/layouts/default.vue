<script setup lang="ts">
const colorMode = useColorMode();

const toggleColorMode = () => {
  colorMode.preference = colorMode.value === "dark" ? "light" : "dark";
};
</script>

<template>
  <ALayoutHeader class=" flex justify-center items-center">
  </ALayoutHeader>
  <ALayoutContent>
    <slot></slot>
  </ALayoutContent>
  <ALayoutFooter class="h-48 flex justify-center items-center">
    <div class="flex gap-2 items-center text-xl">
      <ClientOnly>
        <div @click="toggleColorMode()" class="cursor-pointer">
          <Icon
              :name="$colorMode.value === 'dark' ? 'tabler:moon' : 'tabler:sun'"
          />
        </div>
      </ClientOnly>
    </div>
  </ALayoutFooter>
</template>

<style scoped>

</style>


<script setup lang="ts">

import { EnumErrors } from '@/helpers/enums/Errors.enum';
import {withDefaults, defineProps, ref,watch } from 'vue';

const props = withDefaults(
          defineProps<
            {
              Query:string
            }
          >(),
          {
            Query:""
          }
)

const query = ref(props.Query);

const emit = defineEmits(['OnQuery','OnError'])

watch( () => props.Query, (newQuery)=> query.value = newQuery)

function search(){
  let notValidCharacter:string = ""
  let validateRegex = /[[<>|!#$%^~`'";{}]/g
  const matches = Array.from(query.value.matchAll(validateRegex)).map( mach => mach[0])
  
  notValidCharacter = matches.join("")
  if(notValidCharacter.length > 0) 
    emit('OnError',EnumErrors.NOT_VALID_CHARACTERS + ": " + notValidCharacter)
  else
    emit('OnQuery',query.value)
}
</script>

<template>
  <div class="md:flex md:items-center md:justify-between" >
    <div class="relative flex items-center ">
        <span class="absolute right-0  border-l" >
            <svg xmlns="http://www.w3.org/2000/svg"  @click="search" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor"
              class=" cursor-pointer w-5 h-5 mx-3 text-gray-400 hover:text-purple-500 hover:scale-110   ">
                <path stroke-linecap="round" stroke-linejoin="round" d="M21 21l-5.197-5.197m0 0A7.5 7.5 0 105.196 5.196a7.5 7.5 0 0010.607 10.607z" />
            </svg>
        </span>

        <input
        v-model="query" @keyup.enter="search"
        type="text" placeholder="Search" class="block w-full py-1.5 text-gray-700 bg-white border border-gray-300 rounded-lg md:w-80 placeholder-gray-400/70 px-10     focus:border-transparent  focus:ring-purple-700 focus:outline-none  focus:ring-2 focus:ring-opacity-40">
    </div>
  </div>

</template>


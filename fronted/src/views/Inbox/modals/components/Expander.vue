
<script setup lang="ts">

  import { ref } from 'vue';

  const   isExpanded = ref(true);

</script>
<template>
 
    <div class="content">
        <div class=" space-y-2">
            <div
              class="group flex flex-col gap-2 rounded-lg text-gray-800"
              tabindex="1"
            >
              <div class="  items-center justify-end togle-expand relative"  >
                <span class="bg-white text-gray-500 border-gray-400 hover:text-gray-800 hover:border-gray-800
                absolute right-5  
                 w-8 border rounded-[100%] z-10
                -translate-y-3  text-center">
                  <i @click="isExpanded = !isExpanded"
                    :class="{'-rotate-180':isExpanded}"
                    class=" 
                    text-xl cursor-pointer
                    transition-all fa fa-chevron-down 
                    
                    "
                    aria-hidden="true"
                    ></i>
                </span>
              
              
              </div>
              <div v-if="isExpanded"
                :class="{'visible opacity-100':isExpanded,'invisible h-auto max-h-0  opacity-0':!isExpanded}"
                class="items-center mt-5 border  border-solid  border-gray-600/20  rounded-[5px]
                transition-all  duration-400"
              >
                <slot></slot>
              </div>
            </div>

  


      </div>
    </div>

</template>


<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

.togle-expand::after{
  position: absolute;
  content: "";
  width: 100%;
  height: 10px;
  /* background-color: rgb(221, 219, 219); */
  border-width: 2px 0px 0px 0px;
  border-style: solid;
  border-color: rgba(228, 227, 227, 0.323);
}

</style>

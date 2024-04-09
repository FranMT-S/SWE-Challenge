
<script setup lang="ts">


  import { defineProps, ref ,watch} from 'vue';
  import moment from 'moment'

  import { SubMail } from '@/interfaces/Mail.Interface';
  import MailDetails from './MailDetails.vue'

  const props = defineProps<{
                    mailData:SubMail
                }>()


  const overlayPanel = ref();
  const mail = ref(props.mailData)

  watch(() => props.mailData, (newValue) =>{
    mail.value = newValue
  })

  const toggle = (event) => {
    overlayPanel.value.toggle(event);
  }

 

</script>


<template>
  
  <section class="header my-2" v-if="mail">
      <h1 class="truncate">
        {{mail.Subject}}
      </h1>
      <div class="flex justify-between items-center mt-2 ">
        <div class="truncate">
          <span class="font-semibold">{{mail.From}}. </span> 
          <span class="text-gray-500 font text-xs ">{{ mail.X_From }}</span>
        </div>
        <div class="text-gray-600 text-sm text-center">
          {{ moment(mail.Date).format("DD MMM YYYY, HH:MM:SS A") }}
        </div>
      </div>  
      <div class="text-gray-500 text-sm flex items-center gap-2 ">
        <div class="truncate">
          To {{ mail.To }}
        </div>
        <i class="fa fa-angle-down hover:text-gray-800 hover:scale-110 cursor-pointer" aria-hidden="true"  @click="toggle"></i>
        <OverlayPanel  ref="overlayPanel" class="max-w-2xl max-h-96  py-5 overflow-auto">
          <MailDetails :mail-data="mail"></MailDetails>
        </OverlayPanel>
      </div>  
    </section>
</template>


<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>


</style>

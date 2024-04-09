

<script setup lang="ts">

  import { ref,onBeforeMount } from 'vue';
  import { VueFinalModal } from 'vue-final-modal'

  import { Mail,SubMail } from '../../../interfaces/Mail.Interface';
  import { MailService } from '../../../services/mail.service';

  import Expander from './components/Expander.vue'
  import MailBody from './components/MailBody.vue'
  import MailHeader from './components/MailHeader.vue'


  const props = defineProps<{
    id:string
  }>()





  const mailData = ref({} as Mail)
  const subMessageContent = ref([] as string[])
  const content = ref("")
  const submails = ref([] as SubMail[])
  const firsMessage = ref("")


  const fromRegex = /\n*From:(?<content>[\t \w\d&,.-_]+)/i
  const DateRegex = /\n*Sent:(?<content>[\t \w\d,.-_]+)/i
  const ToRegex = /\n*(To|Re):(?<content>[\t \w\d&,.-_]+)/i
  const SubjectRegex = /\n*Subject:(?<content>[\t \w\d&,.-_]+)/i
  const CCRegex = /\n*CC:(?<content>[\t \w\d&,.-_]+)/i
  const messageRegex : RegExp = /-----Original Message-----/g;




  onBeforeMount(  () =>{
  
    MailService.getMail(props.id)
      .then( res => res.data.data)
      .then( mail =>{
        mailData.value = mail

        content.value = mail.Content;
        subMessageContent.value =  content.value.split(messageRegex)
        firsMessage.value = subMessageContent.value[0]
      
    
        for (let i = 1; i < subMessageContent.value.length; i++) {
          let _subMail:SubMail = { From: "", To:"", Subject: "", Date: new Date()}
    
        
          let match = fromRegex.exec(subMessageContent.value[i])
          if(match){
            subMessageContent.value[i] = subMessageContent.value[i].replace(fromRegex,"")
            _subMail.From = match.groups!['content'].trim() || ""
          }
      
          match = DateRegex.exec(subMessageContent.value[i])
          if(match){
            subMessageContent.value[i] = subMessageContent.value[i].replace(DateRegex,"")
            _subMail.Date = new Date(match.groups!['content'].trim()) || ""
          }

          match = ToRegex.exec(subMessageContent.value[i])
          if(match){
            subMessageContent.value[i] = subMessageContent.value[i].replace(ToRegex,"")
            _subMail.To =match.groups!['content'].trim() || ""
          }

          match = CCRegex.exec(subMessageContent.value[i])
          if(match){
            subMessageContent.value[i] = subMessageContent.value[i].replace(CCRegex,"")
            _subMail.Cc =match.groups!['content'].trim() || ""
          }

          match = SubjectRegex.exec(subMessageContent.value[i])
          if(match){
            subMessageContent.value[i] = subMessageContent.value[i].replace(SubjectRegex,"")
            _subMail.Subject =match.groups!['content'].trim() || ""
          }

          _subMail.Content = subMessageContent.value[i].trim()

    
          submails.value.push(_subMail)
        }

        
      })
      .catch(console.log)
  })

</script>


<template>
  <VueFinalModal
        class="confirm-modal "
        content-class="confirm-modal-content justify-center   absolute inset-0 m-auto h-5/6  w-4/6"
        overlay-transition="vfm-fade-up"
        content-transition="vfm-slide-down"
  >
    <section class="h-full content-center"> 
      <section class="overflow-auto h-full" >
        <div class="mx-3 h-full ">
          <div class="rounded-[5px]  mt-5 h-full" >
            <section >
              <div v-if="mailData"
                class="items-center mb-5 border  
                transition-all  duration-400"
              >
            
              <MailHeader class="px-12" :mail-data="mailData"></MailHeader>
              <hr>
              <MailBody class="px-12 my-4 "  v-if="subMessageContent[0]" :content="subMessageContent[0]"></MailBody>
              
              </div>
              <Expander class="mb-10" v-for="(submail, index) in submails" :key="index">
                <MailHeader class="px-12" :mail-data="submail"></MailHeader>
                <hr>
                <MailBody class="px-12 my-4 break-words" :content="submail.Content"></MailBody>
              </Expander>
            
            </section>
          </div>
        </div>  
      </section>
    </section>
  </VueFinalModal>
</template>

<style>

  .mail_header{
    width: 66.666667%
  }

  .mail_header_field{
    flex-basis:6.0rem;
    text-align:right;

  }

  .mail_header_data{
    margin-left: 0.5rem;
  }

 .confirm-modal {
    display: flex;
    justify-content: center;
    align-items: center;
  }
  .confirm-modal-content {
    display: flex;
    flex-direction: column;
    padding: 1rem;
    background: #fff;
    border-radius: 0.5rem;
  }
  .confirm-modal-content > * + *{
    margin: 0.5rem 0;
  }
  .confirm-modal-content h1 {
    font-size: 1.375rem;
  }

  .dark .confirm-modal-content {
    background: #000;
  } 

</style>
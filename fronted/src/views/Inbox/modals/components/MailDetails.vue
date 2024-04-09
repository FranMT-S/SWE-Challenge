
<script setup lang="ts">

  import { defineProps, ref,onBeforeMount} from 'vue';
  import moment from 'moment'

  import { SubMail } from '@/interfaces/Mail.Interface';
  
  const props = defineProps<{
                  mailData:SubMail
                }>()

  const ToFieldZip = ref([] as string[][])
  const bccFieldZip = ref([] as string[][])
  const ccFieldZip = ref([] as string[][])
  const mail = ref(props.mailData)




  onBeforeMount(() => {
    let X_FieldRegex = /(?<=>),|(?<=')\s/
    if(mail.value.To && mail.value.X_To){
      ToFieldZip.value = zip(mail.value.To?.split(','),mail.value.X_To.split(X_FieldRegex) )
    }

    if(mail.value.Bcc && mail.value.X_bcc){
      bccFieldZip.value = zip(mail.value.Bcc?.split(','),mail.value.X_bcc.split(X_FieldRegex) )
    }
  
    if(mail.value.Cc && mail.value.X_cc){
      ccFieldZip.value = zip(mail.value.Cc?.split(','),mail.value.X_cc.split(X_FieldRegex) )
    }
  })
  
  const zip = (a:string[], b:string[]) => {
    if(a.length == b.length){
      return a.map((k, i) => [k, b[i].replaceAll(/^'|'$/g,"")])
    }

    return []
  };


</script>


<template>
  <section  v-if="mail">
    <div class="details">
      <section class="details_container break-all grid grid-cols-[span 1 / 1fr] grid-flow-row-dense gap-y-2 gap-x-3 whitespace-break-spaces break-before-all">
        <div class="details" v-if="mail.From">
          <span class="details_field">From:</span>
          <span class="details_data font-semibold">
            {{mail.From}} 
              <span class="text-gray-500 font-normal text-sm">{{mail.X_From}}</span>
          </span>
        </div>

        <div class="details" v-if="mail.To || mail.X_To">
          <span class="details_field">To:</span>
          <span class="details_data" >
            <span v-if="ToFieldZip.length == 0">
              {{mail.To}}
            </span>
            <span v-else v-for="(To,index) in ToFieldZip">
              {{To[0]}}
              <span class="text-gray-500 font-normal text-sm">{{To[1]}}</span>
              <span v-if="index < ToFieldZip.length - 1">,</span>
            </span>
          </span>
        </div>
    
        <div class="details" v-if="mail.Date">
          <span class="details_field">Date:</span>
          <span class="details_data">{{moment(mail.Date).format("DD MMM YYYY, HH:MM:SS A")}}</span>
        </div>
    
        <div class="details" v-if="mail.Subject">
          <span class="details_field">Subject:</span>
          <span class="details_data">{{mail.Subject}}</span>
        </div>

        <div class="details" v-if="mail.Bcc || mail.X_bcc">
          <span class="details_field">BCC:</span>
          <span class="details_data" >
            <span v-if="bccFieldZip.length == 0">
              {{mail.Bcc}}
            </span>
            <span v-else v-for="(bcc,index) in bccFieldZip">
              {{bcc[0]}}
              <span class="text-gray-500 font-normal text-sm">{{bcc[1]}}</span>
              <span v-if="index < bccFieldZip.length - 1">,</span>
            </span>
          </span>
        </div>


        <div class="details" v-if="mail.Cc || mail.X_cc">
          <span class="details_field">CC:</span>
          <span class="details_data" >
            <span v-if="ccFieldZip.length == 0">
              {{mail.Cc}}
            </span>
            <span v-else v-for="(cc,index) in ccFieldZip">
              {{cc[0]}}
              <span class="text-gray-500 font-normal text-sm">{{cc[1]}}</span>
              <span v-if="index < ccFieldZip.length - 1">,</span>
            </span>
          </span>
        </div>

        <div class="details" v-if="mail.X_FileName">
          <span class="details_field">FileName:</span>
          <span class="details_data">{{mail.X_FileName}}</span>
        </div>
        
        <div class="details" v-if="mail.X_Folder">
          <span class="details_field">Folder:</span>
          <span class="details_data">{{mail.X_Folder}}</span>
        </div>
      </section>
    </div>
  </section>
</template>


<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

.details{
  display: contents;
}
.details_field{
  grid-column-start: 1 ;
  grid-row: span 1;
  color:gray;
  text-align: right;
}
.details_data{
  grid-column-start: 2 ;
  grid-row: span 1;
  word-wrap: break-word;
  color:black
}


</style>

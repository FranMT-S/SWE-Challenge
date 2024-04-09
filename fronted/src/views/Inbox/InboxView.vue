
<script setup lang="ts">

import {  useRoute, useRouter } from 'vue-router'
import { ref,onBeforeMount, watch,computed} from 'vue';

import { ERouterName } from '@/helpers/enums/RouterName.enum';
import { MailService } from '@/services/mail.service';
import {EnumErrors} from '@/helpers/enums/Errors.enum'
import { Hit, ResponseData, ResponseHits } from '@/interfaces/Mail.Interface';

import MailViewModal from './modals/MailViewModal.vue'
import MyTable from '@/components/Table.vue';
import Pagination from '@/components/Pagination.vue';
import SearchBar from '@/components/SearchBar.vue';

import { ModalsContainer, useModal } from 'vue-final-modal'
import { AxiosResponse } from 'axios';



const fields =  ["From","To","Subject","Cc","Bcc","X_Folder","X_Origin","X_FileName","Content"]

const router = useRouter();
const route = useRoute();

const overlayPanel = ref();
const query = ref("");
const containtSome = ref("");
const notContaint = ref("");
const containtAll = ref("");
const fieldSelect = ref("");
const fieldValue = ref("");
const activedAdvanceSearch = ref(false);
const errMsg = ref("");
const max = 100;
const page = ref(0)
const total = ref(0)
const Hits = ref([] as Hit[])
var debounceTimeout = 0

onBeforeMount(() => {
  page.value = Number(route.query['page']?.toString()) || 1;
  page.value = page.value < 1 ? 1 : page.value;

  query.value = route.query['query']?.toString()  || "" ;

  getMails(query.value,(page.value - 1) * max,max)
})



const getMails = (query:string,from:number,max:number) => {

  if (from < 0 ) from = 0;
  if(query.trim().length == 0)
    getAllMails(from,max)
  else
    findMails(query,from,max)
  
}

 const getAllMails = (from:number,max:number) => {
    MailService.getAllMails(from,max).then(responseHits)
  .then( hits => Hits.value = hits )
  .catch(catchError)
}

const findMails = (query:string,from:number,max:number) => {
    MailService.findMails(query,from,max).then(responseHits)
  .then( hits => Hits.value = hits )
  .catch(catchError)
}

const responseHits = (response:AxiosResponse<ResponseData<ResponseHits>, any>) => {
    errMsg.value = ""
  
    total.value = response.data.data.total.value as number
    if(total.value == 0 ) page.value = 0
    return response.data.data.hits as Hit[] 
}

const catchError = (error) => {
  
    if(error.code == "ERR_NETWORK"){
      errMsg.value = EnumErrors.CONNECTION
      console.log(error)
    }else{
      errMsg.value = error.response.data.msg as string
      console.error(error);
    }

}

const toggle = (event) => {
  overlayPanel.value.toggle(event);
}


// when params change get mails
watch(
      () => route.params,
      () => {  
        // This is in case you change pages very quickly
        clearTimeout(debounceTimeout);
        debounceTimeout = setTimeout(() => {
            getMails(query.value,(page.value - 1) * max,max)
          }, 150); 
        },       
)

watch(
  query,
  (newQuery) => {
    router.push({name:ERouterName.MAILS,query:{query:newQuery,page:page.value}})
  }
)


watch(
  page,
  (newPage) => {
    router.push({name:ERouterName.MAILS,query:{query:query.value,page:newPage}})
  }
)

const changeQuery = (newQuery) => {
  query.value = newQuery;
  page.value = 1;
}

const changePage = (newPage) => {
  if(newPage < 1) newPage = 1
   page.value = newPage
  if(newPage > lastPage.value)
    page.value = lastPage.value

}

const advanceSearch = () =>{
  let newQuery = ""
  newQuery += containtSome.value.trim()
  if(notContaint.value)
    notContaint.value.trim().split(/\s+/).forEach( word => newQuery += ` -${word}`)
  if(containtAll.value)
    containtAll.value.trim().split(/\s+/).forEach( word => newQuery += ` +${word}`)
  if(fieldValue.value && fieldSelect.value)
    newQuery += ` ${fieldSelect.value}:${fieldValue.value.trim()}`
  
  query.value = newQuery

  containtSome.value = ""
  notContaint.value = ""
  containtAll.value = ""
  fieldValue.value = ""
  fieldSelect.value = ""
  activedAdvanceSearch.value = false
}

const verifyKey = (event) =>{
  var key = event.key;
  var keyBlock =`.,:";'"()[]{}\`<>''!#$%^&*?+*-/\\=|`;
  if (keyBlock.includes(key)) {  
      event.preventDefault();
  }
}

const openModal = (id) => {
  const { open, close } = useModal({
  component: MailViewModal,
  attrs: {
    id: id,
    onConfirm() {
      close()
    },
  },
  })

  open()
}




const lastPage = computed(() => {
  return total.value == 0 ? 1 : total.value % max == 0 ? total.value / max : Math.ceil(total.value / max)
})

const messageTotalStart = computed(() => {
  return (page.value - 1) * max < total.value ? (page.value - 1) * max + 1 : total
})

const messageTotalEnd = computed(() => {
  return (page.value) * max < total.value ? (page.value) * max : total
})

const IsInvalidAdvancedSearch = computed(() => {
  return !(
      containtSome.value.length > 0 || 
      notContaint.value.length > 0 || 
      containtAll.value.length > 0 || 
      (fieldValue.value.length > 0 && fieldSelect.value.length > 0 )
  )
})

</script>



<template>
  <div class="inbox">
    <Message class="m-auto p-component p-message p-message-error w-4/6"   @close="errMsg = '' " v-if="errMsg?.length > 0" severity="error">{{errMsg}}</Message>

    <div class=" text-center  divide-y  rounded ">
      <div class="flex justify-evenly items-end my-2 py-3 flex-wrap-reverse   w-5/6">
        <div class=" ">       
          <div class="flex items-center gap-3">          
            <SearchBar :Query="query" @OnQuery="changeQuery" @OnError="(event:any) => errMsg = event "></SearchBar>
            <span class="relative" v-tooltip="{ value: 'Advanced search', showDelay: 400 }">   
              <Button @click="toggle"  icon="fa fa-ellipsis-vertical" severity="help" text raised  unstyled  rounded aria-label="Favorite" 
              :class="{'bg-gray-50 scale-110 text-purple-800':activedAdvanceSearch}"
              class="w-12 h-12 text-center   hover:scale-110 text-gray-400 hover:text-purple-800 hover:bg-gray-50 rounded-[100%]"
              />
              <OverlayPanel @show="activedAdvanceSearch = true" @hide="activedAdvanceSearch = false"  ref="overlayPanel" class="absolute">           
                <form @submit.prevent="advanceSearch" >
     
                  <div class=" p-5 grid-flow-col-dense flex flex-col gap-3  bg-white z-10 " >
                    <input type="text" placeholder="Contain some" @keydown="verifyKey"  v-model="containtSome"
                      class="text-center  border border-gray-400 rounded  focus:border-transparent  focus:ring-purple-700 focus:outline-none  focus:ring-2 focus:ring-opacity-40" >
                  
                  
                    <input type="text"  placeholder="Not contain"  @keydown="verifyKey" v-model="notContaint"
                      class="text-center border border-gray-400 rounded  focus:border-transparent  focus:ring-purple-700 focus:outline-none  focus:ring-2 focus:ring-opacity-40" >
                  

                    <input type="text"  placeholder="Contain all"  @keydown="verifyKey" v-model="containtAll"
                      class="text-center border border-gray-400 rounded  focus:border-transparent  focus:ring-purple-700 focus:outline-none  focus:ring-2 focus:ring-opacity-40" >
                  
                    <div class="flex gap-2 flex-wrap md:flex-nowrap  ">
                        <select  class="flex-grow md:flex-grow-0  border border-gray-400 rounded "  v-model="fieldSelect">
                          <option value="" class="text-center" >Select Field</option>
                          <option class="text-center" v-for="(field, index) in fields" :key="index" :value="field">{{ field.replace('_',"-") }}</option>
                      </select>
                      <input type="text"  placeholder="Contain"  @keydown="verifyKey" v-model="fieldValue"
                        class="text-center border border-gray-400 rounded  focus:border-transparent  focus:ring-purple-700 focus:outline-none  focus:ring-2 focus:ring-opacity-40" >        
                    </div>
                  
                  
                    <div class="col-span-full flex justify-center">
                    <button type="submit" :disabled="IsInvalidAdvancedSearch" 
                    class="px-12 mt-5 border disabled:opacity-50 disabled:pointer-events-none border-purple-800 text-purple-500 rounded-md hover:bg-purple-800 hover:text-white">
                    <i class="fa fa-search" aria-hidden="true"></i>  Search
                    </button>
                    </div>
                  </div>
                </form>
              </OverlayPanel>                          
            </span>   
          </div>
        </div>
        <div class="mb-4 md:mb-0">
  
        
        </div>
      </div>
      <div >
        <div class="container px-4 flex md:justify-end  gap-x-16 mx-5 pt-2 flex-wrap ">
          <div>
            <i class="fa-regular fa-envelope" v-tooltip.left="{ value: 'Total messages', showDelay: 300 }"></i> {{ messageTotalStart }}-{{ messageTotalEnd }} / {{ total }}
          </div>
          <div>
            <i class="fa-regular fa-file" v-tooltip.left="{ value: 'Page', showDelay: 300 }"></i> {{ page }} / {{ lastPage }}
          </div>
        </div>
        <MyTable :working="false" @on-id-mail="openModal"  :query="query" class="mb-10" :Hits="Hits"></MyTable>
        <Pagination :Max="lastPage" :Page="page" :total="total" @on-page="changePage"></Pagination>
   
        <ModalsContainer />
      </div>
    </div>
  </div>
</template>

<style scoped>
.p-message .p-message-wrapper{ justify-content: center !important;}
</style>
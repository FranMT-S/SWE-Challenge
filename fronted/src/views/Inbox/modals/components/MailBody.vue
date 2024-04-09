
<script setup lang="ts">



import {withDefaults, defineProps, ref,onBeforeMount } from 'vue';


  const props = withDefaults(
    defineProps<
      {
        content:string
      }
    >(),
    {
      content:""
    }
  )

  const content = ref(props.content);

  const fields = ['From:','Sent:','To:','Subject:','mailto:','E?mail:',"Phone:","Fax:","CC:"]

  const mailRegex : RegExp = new RegExp(`([a-zA-Z0-9._-]+@[a-zA-Z0-9._-]+\.[a-zA-Z0-9._-]+)`,'gi')
  const FieldRegex : RegExp = new RegExp(`(${fields.join("|")})`,'gi');
  const FileRegex  = /<<\s*File:(\s*[\w+,_\s-()\d]+)\.(\w+)\s*>>/gi;
  const listRegex  = /\n\s*- ([\w\d\. -]+)/gi;
  const ForwardedRegex  = /-+\s+(Forwarded by[A-z-0-9\s\/:_()]+)\s+-+/gi;

  onBeforeMount( () =>{

    content.value = content.value.replaceAll(FieldRegex,"<span class='font-semibold'>$1</span>")
    content.value = content.value.replaceAll(mailRegex,"<a href=$1 class='text-blue-800'>$1</a>")
    content.value = content.value.replaceAll(ForwardedRegex,"<p class='text-gray-500 text-xs whitespace-normal' >$1</p>")

    content.value = content.value.replaceAll(FileRegex,`
    <div class="group/file border flex flex-col p-4 rounded justify-center   w-min  items-center cursor-pointer hover:bg-gray-100 file">
        <i class="fa fa-file text-blue-500" aria-hidden="true"></i>     
        <span class="text-gray-400 group-hover/file:text-gray-800">$1.$2</span>
    </div>
    `)
    content.value = content.value.replaceAll(listRegex,"<li>$1</li>")
  })

</script>


<template>
   <div class="whitespace-pre-wrap break-words" v-html="content"> </div>
</template>


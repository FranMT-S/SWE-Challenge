
<script setup lang="ts">

import { withDefaults, defineProps,ref,watch,computed} from 'vue';


const props = withDefaults(
    defineProps<
        {
            Page:number,
            Max:number,
            total:number
        }
    >(),
    {
        Page : 1
    }
)

const emit =  defineEmits<{
    (e:'OnPage',newPage:number):void
}>()

const page = ref(props.Page)
const max = ref(props.Max)


function changePage(newPage:number){
    page.value = newPage
}

watch(() => props.Page,(newPage) =>{
    
    page.value = newPage
})

watch(() => props.Max,(newPage) =>{
    max.value = newPage
})

watch(page,(newPage) =>{
    emit('OnPage',newPage)
})


</script>


<template>
    <div class="card">
        <Paginator class=" py-3"
            @page="({page}) =>changePage(page + 1)"
            :first="(page - 1) * 100"

            :template="{
                '640px': ' PrevPageLink JumpToPageDropdown NextPageLink ',
                '960px': 'FirstPageLink PrevPageLink JumpToPageDropdown NextPageLink LastPageLink',
                '1300px': 'FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink',
                default: 'FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink JumpToPageDropdown '
            }"
            :rows="100"
            :totalRecords="total"
            :pt="{
            

                root: {
                    class: [
                        'flex items-center justify-center flex-wrap',
                        'bg-white'
                    ],
                    
                },
                firstpagebutton: ({ context }) => ({
                    class: [
                        'relative inline-flex items-center justify-center user-none overflow-hidden leading-none',
                        'border-0 text-gray-500  min-w-[3rem] h-12 m-[0.143rem] ',
                        'transition duration-200',
                        'text-gray-500 hover:text-black hover:bg-gray-100', //Dark Mode
                        {
                            'cursor-default pointer-events-none opacity-60': context.disabled,
                        }
                    ]
                }),
                previouspagebutton: ({ context }) => ({
                    class: [
                        'relative inline-flex items-center justify-center user-none overflow-hidden leading-none',
                        'border-0 text-gray-500 min-w-[3rem] h-12 m-[0.143rem] ',
                        'transition duration-200',
                        'text-gray-500 hover:text-black hover:bg-gray-100', //Dark Mode
                        {
                            'cursor-default pointer-events-none opacity-60': context.disabled,
                        }
                    ]
                }),
                nextpagebutton: ({ context }) => ({
                    class: [
                        'relative inline-flex items-center justify-center user-none overflow-hidden leading-none',
                        'border-0 text-gray-500 min-w-[3rem] h-12 m-[0.143rem] ',
                        'transition duration-200',
                        'text-gray-500 hover:text-black hover:bg-gray-100', //Dark Mode
                        {
                            'cursor-default pointer-events-none opacity-60': context.disabled,
                        }
                    ]
                }),
                lastpagebutton: ({ context }) => ({
                    class: [
                        'relative inline-flex items-center justify-center user-none overflow-hidden leading-none',
                        'border-0 text-gray-500 min-w-[3rem] h-12 m-[0.143rem] ',
                        'transition duration-200',
                        'text-gray-500 hover:text-black hover:bg-gray-100', //Dark Mode
                        {
                            'cursor-default pointer-events-none opacity-60': context.disabled,
                        }
                    ]
                }),
                pagebutton: ({ context }) => ({
                    class: [
                        'relative inline-flex items-center justify-center user-none overflow-hidden leading-none',
                        'border-0 text-gray-500 min-w-[3rem] h-12 m-[0.143rem] ',
                        {
                            'bg-purple-800 border-blue-50 text-blue-700 text-white font-semibold': context.active,
                            'hover:bg-gray-100 hover:text-gray-800':!context.active
                        }
                    ]
                }),
                jumptopagedropdown: {
                    root: ({ props, state }) => ({
                        class: [
                            'inline-flex relative cursor-pointer user-none',
                            'bg-white border rounded-md',
                            'transition duration-100',
                            'h-12 mx-2',
                            {
                                'outline-none outline-offset-0 shadow-[0_0_0_0.2rem_rgba(91,41,148,0.2)] border-purple-500': state.focused && !props.disabled, //Focus
                                'border-gray-300': !state.focused,
                                'hover:border-purple-500': !props.disabled //Hover
                            }
                        ]
                    }),
                    input: {
                        class: [
                            'font-sans text-base text-gray-600 p-3 m-0 rounded-md apperance-none',
                            'block whitespace-nowrap overflow-hidden flex-auto w-[1%] cursor-pointer text-ellipsis border-0 pr-0',
                            'focus:outline-none focus:outline-offset-0',
                            'focus:text-gray-800'
                        ]
                    },
                    trigger:({ props, state }) => ({
                        class: [
                            'flex items-center justify-center shrink-0', 
                            'text-gray-500  w-12 rounded-r-md',
                            'hover:text-black-800',
                            {'text-gray-800':state.focused},
                           
                        ]
                    }),
                    panel: {
                        class: [
                      
                        ]
                    },
                    wrapper: 'overflow-auto',
                    list: 'm-0 p-0 py-3 list-none',
                    item: ({ context }) => ({
                        class: [
                            'relative font-normal cursor-pointer space-nowrap overflow-hidden',
                            'm-0 py-3 px-5 border-none text-gray-600 rounded-none',
                            'transition duration-200',
                            {
                                'bg-purple-300/40 text-purple-500': context.focused && context.selected,
                            }
                        ]
                    })
                }
                
            }">
            
        </Paginator>
    </div>
    
</template>



<style scoped>


</style>

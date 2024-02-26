<template>
    <div 
        :class="'top-click-detect popup-selector-main overlay-' + this.ledgerType + (this.isLifetime ? '-lifetime' : '')"
        v-on:click="e => clickTop(e)"
    >
        <div class="inner-click-detect popup-selector-inner">
            <button class="detect-trigger close-button" v-on:click="e => { e.preventDefault(); $emit('close')}">
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                </svg>
            </button>
            <input
                :id="(ledgerType) + '-search' + (isLifetime ? '-lifetime' : '')"
                class="input-search" type="text" placeholder="Search..."
                v-on:focus="e => {e.preventDefault(); e.stopPropagation(); }"
                v-on:input="e => {e.preventDefault(); $emit('input', e.target.value)}"
              />
              <div class="popup-opt-container">
                <div v-for="item in itemList" v-bind:key="item" class="popup-opt"
                    v-on:click="e => {e.preventDefault(); $emit('select', item)}"
                >
                    <img 
                        class="mr-1rem max-w-7" v-if="item.imagePath != null && item.imagePath != ''"
                        :alt="item.text" :src="getImgPath(item)"
                        :style="(ledgerType === 'drop' && item.rarityGif ? 'background: url(' + item.rarityGif + ') center center no-repeat; background-size: cover;' : '')"
                    >
                    <span :class="(ledgerType === 'drop' ? (item.styleClass ?? '') : 'text-gray-400')">
                        {{ item.text }}
                    </span>
                </div>
                <div v-if="itemList == [] || itemList.length == 0">
                  <span class="text-gray-400">
                    No results found.
                  </span>
                </div>
              </div>
        </div>
    </div>
</template>

<script>
    export default {
        emits: ['close', 'input', 'select'],
        props: {
            itemList: Array,
            ledgerType: String,
            isLifetime: Boolean,
        },
        methods: {
            getImgPath(item){
                if(this.ledgerType === 'drop') return `images/${item.imagePath}`;
                else if(this.ledgerType === 'target') return `images/targets/${item.imagePath}`;
            },
            clickTop(e){
                const topElement = e.target;
                const innerEl = topElement.querySelector('.inner-click-detect');
                if(innerEl === null) return;
                if(innerEl?.classList?.contains('hidden')) return;
                  const divRect = innerEl.getBoundingClientRect();
                  if(e.clientX < divRect.left || e.clientX > divRect.right || e.clientY < divRect.top || e.clientY > divRect.bottom)
                    innerEl.querySelector('.detect-trigger').click();
            }
        }
    }
</script>
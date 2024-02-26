<template>
    <div v-if="ifCount > 0">
        <div :class="labelClassList">
            {{ labelDisplayValue }} ({{ ifCount.toString() }})
        </div>
        <div :class="getRepeatClass()">
            <div :class="getInnerRepeatClass()" v-for="(item, itemIndex) in itemArray">
                <a v-external-link :class="getAClass(item)"
                    target="_blank" :href="afExplorerName(item, getSpecPathOffset())">
                    <img class="h-full w-full" :alt="item.gameName" :src="specPath(item, getSpecPathOffset())"/>
                    <div v-if="item.count > 1" :class="'ledger-af-count ' + this.numToDigClass(item.count)">
                        {{ item.count }}
                    </div>
                    <span class="text-sm tooltiptext-custom speech-bubble">
                        {{ item.gameName }} (T{{parseInt(item.level) + parseInt(getSpecPathOffset())}}) 
                        <span v-if="type == 'artifact'" :class="afRarityClass(item)">
                          {{ afRarityText(item) }}
                        </span>
                        × {{ item.count }}
                        <br v-if="type == 'artifact' || type == 'stone'"/>
                        <span v-if="type == 'artifact' || type == 'stone'">
                            {{ boundEffectString(item.effectString)[0] }}
                            <span class="text-green-500">
                                {{ boundEffectString(item.effectString)[1] }}
                            </span>
                            {{ boundEffectString(item.effectString)[2] }}
                        </span>
                        <hr v-if="ledgerType == 'lifetime' && lifetimeShowPerShip" class="mt-05rem mb-05rem w-full">
                        <span v-if="ledgerType == 'lifetime' && lifetimeShowPerShip">
                            (<span class="text-green-500">{{ (item.count / lifetimeMissionCount).toFixed(5) }}</span> per ship - 
                            <span class="text-green-500">1</span>:<span class="text-green-500">{{ (1 / (item.count / lifetimeMissionCount)).toFixed(2) }}</span>)
                        </span>
                    </span>
                </a>
            </div>
        </div>
        <br />
    </div>
</template>

<script>
    export default {
        props: {
            useGifsForRarity: Boolean,
            labelClassList: String,
            labelDisplayValue: String,
            ifCount: Number,
            itemArray: Array,
            type: String,
            ledgerType: String,
            lifetimeShowPerShip: Boolean,
            lifetimeMissionCount: Number,
            afRarityClass: Function,
            afRarityText: Function,
        },
        methods: {
            getRepeatClass(){
                if(this.ledgerType === 'lifetime') return 'ledger-af-repeat-lifetime';
                else return 'ledger-af-repeat';
            },
            getInnerRepeatClass(){
                if(this.ledgerType === 'lifetime') return 'af-view-rep-lifetime';
                else return 'mission-view-rep';
            },
            getSpecPathOffset(){
                if(this.type == 'stone') return '2';
                else return '1';
            },
            afBackgroundClass(item){
              switch (item.rarity) {
                case 0: return 'bg-common';
                case 1: return (this.useGifsForRarity ? "bg-cover bg-rare-animated" : 'bg-rare');
                case 2: return (this.useGifsForRarity ? "bg-cover bg-epic-animated" : 'bg-epic')
                case 3: return (this.useGifsForRarity ? "bg-cover bg-legendary-animated" : 'bg-legendary');
                default: return 'bg-common';
              }
            },
            getAClass(af){
                if(this.type == 'artifact') return 'ledger-af-link tooltip-custom ' + this.afBackgroundClass(af);
                else return 'ledger-af-link tooltip-custom bg-common';
            },
            afExplorerName(drop, addend) {
              return 'https://wasmegg-carpet.netlify.app/artifact-explorer/#/artifact/' + 
                drop.name.replace('_FRAGMENT', '').toLowerCase().replace("_", "-") + '-' + (parseInt(drop.level) + parseInt(addend));
            },
            specPath(spec, addend) {
              const fixedName = spec.name.replaceAll("_FRAGMENT", "").replaceAll("ORNATE_GUSSET", "GUSSET").replaceAll("VIAL_MARTIAN_DUST", "VIAL_OF_MARTIAN_DUST");
              return "images/artifacts/" + fixedName + "/" + fixedName + "_" + (parseInt(spec.level) + parseInt(addend)) + ".png";
            },
            numToDigClass(num){
              const parsedNum = parseInt(num.toString());
              switch (true) {
                case parsedNum > 99999: return "w-sixdig";
                case parsedNum > 9999: return "w-fivedig";
                case parsedNum > 999: return "w-fourdig";
                case parsedNum > 99: return "w-threedig";
                case parsedNum > 9: return "w-twodig";
                default: return "w-onedig";
              }
            },
            boundEffectString(str) {
              if(str.startsWith('!!')){
                const match = /<([^>]+)>/g.exec(str);
                return (!match ? ["?", "?", "?"] : [str.substring(2, match.index), match[1], ""]);
              }
              const match = /\[(.*?)\]/g.exec(str);
              return (!match ? ["?", "?", "?"] : [str.substring(0, match.index), match[1], str.substring(match.index + match[0].length)]);
            }
        }
    }
</script>
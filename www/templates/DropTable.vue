<template>
    <span v-if="any()" class="py-4 px-5 font-bold text-left flex flex-auto flex-col min-w-70vw max-w-70vw">
        <span :class="'text-base text-rarity-' + rarity">
            {{ getRarityText() }}
            <parened-item :inner-class="'text-base text-rarity-' + rarity" :inner-text="sumByCount()"></parened-item>
            <button type="button" class="toggle-link ml-1rem" v-on:click="itemsExpanded = !itemsExpanded">
                {{ itemsExpanded ? 'Collapse to List' : 'Expand to Table' }}
            </button>
        </span>
        <inner-drop-display
            v-if="!itemsExpanded"
            class="text-center text-gray-400 mt-2"
            :options="{ 'noJustify': true }" :item-array="where()"
            ledger-type="lifetime" :total-drops-count="sumByCount()"
            :lifetime-show-per-ship="false" :show-expected-drops="false"
        ></inner-drop-display>
        <table v-if="itemsExpanded">
            <th><tr><td colspan="2"></td></tr></th>
            <tbody>
                <template v-for="item in where()" :key="item.id">
                    <tr><td colspan="4"><hr class="mt-0_5rem mb-0_5rem w-full"></td></tr>
                    <tr>
                        <td colspan="1" :class="'text-base text-rarity-' + rarity">
                            <div class="flex flex-row items-center">
                                <inner-drop-display
                                    class="text-center text-gray-400 mt-2"
                                    :options="{ 'noJustify': true }" :item-array="[item]"
                                    ledger-type="lifetime" :total-drops-count="sumByCount()"
                                    :lifetime-show-per-ship="false" :show-expected-drops="false"
                                ></inner-drop-display>
                                <span class="ml-1rem">
                                    {{ item.gameName }} <parened-item class="text-gray-400" :inner-class="'text-base text-rarity-' + rarity" :inner-text="item.count"></parened-item>
                                </span>
                            </div>
                        </td>
                        <td colspan="3" class="text-left">
                            <span>
                                <span class="text-green-500">{{ item.missionInfos.length }} Mission(s)</span> 
                                <button type="button" class="toggle-link ml-1rem" v-on:click="item.digShipInfo = !item.digShipInfo">
                                    {{ (item.digShipInfo ? 'Collapse Missions' : 'Expand Missions') }}
                                </button>
                            </span>
                            <div v-if="item.digShipInfo" class="flex flex-col">
                                <div v-for="mission in item.missionInfos">
                                    {{ mission.missionId }}
                                </div>
                            </div>
                        </td>
                    </tr>  
                </template>
            </tbody>
        </table>
    </span>
</template>

<script>
    import InnerDropDisplay from './InnerDropDisplay.vue';
    import ParenedItem from './ParenedItem.vue';

    export default {
        emits: [],
        components: {
            InnerDropDisplay,
            ParenedItem,
        },
        props: {
            data: Object,
            predicate: Function,
            rarity: Number,
            itemsExpanded: Boolean,
        },
        methods: {
            getPredicate(){
                return this.predicate ?? (item => item.rarity == this.rarity);
            },
            any(){
                return this.getItems().some(this.getPredicate());
            },
            where(){
                return this.getItems().filter(this.getPredicate());
            },
            sumByCount(){
                return this.where().reduce((acc, item) => (acc + (item?.count ?? 0)), 0);
            },
            getItems(){
                return [
                    ...this.data.artifacts,
                    ...this.data.stones,
                    ...this.data.ingredients,
                    ...this.data.stoneFragments,
                ];
            },
            getRarityText(){
                switch(parseInt(this.rarity)){
                    case 3: return 'Legendary';
                    case 2: return 'Epic';
                    case 1: return 'Rare';
                    case 0: return 'Common';
                }
            }
        },
    }
</script>
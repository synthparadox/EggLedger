<template>
    <img v-if="targetImageResult !== null"
         :src="targetImageResult[0]"
         :alt="targetImageResult[1]"
         class="target-ico" />
  </template>
  
  <script>
  export default {
    props: {
      target: String
    },
    methods: {
      targetImage(target) {
        if (target == null || target === "" || target === "UNKNOWN") {
          return null; // Return null if target is empty or UNKNOWN
        }
        target = target.toUpperCase()
            .replaceAll("ORNATE_GUSSET", "GUSSET")
            .replaceAll("VIAL_MARTIAN_DUST", "VIAL_OF_MARTIAN_DUST");
        let tier = 4;
        if (target.indexOf('_FRAGMENT') > -1) {
            target = target.replace('_FRAGMENT', '');
            tier = 1;
        }
        if (target === 'GOLD_METEORITE' || target === 'TAU_CETI_GEODE' || target === 'SOLAR_TITANIUM') {
            tier = 3;
        }
        const path = `images/artifacts/${target}/${target}_${tier}.png`;
        return [path, target];
      }
    },
    computed: {
      targetImageResult() {
        return this.targetImage(this.target);
      }
    }
  };
  </script>
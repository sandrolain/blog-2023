document.addEventListener("DOMContentLoaded", () => {
  const tds = Array.from(document.querySelectorAll("td"));
  for(const td of tds) {
    if(td.innerText.indexOf("▢") > -1) {
      const lvl = td.innerText.trim().length;
      td.innerHTML = `<div class="about-lvl" data-lvl="${lvl}"></div>`;
    }
  }

  const pluralsightTitle = document.getElementById("pluralsight");
  if(pluralsightTitle) {
    (async() => {
      const response = await fetch("/pluralsight-badges.json");
      const badges = await response.json();
      let html = [];
      const escaper = document.createElement("div");
      const escape = (text) => {
        escaper.innerText = text;
        return escaper.innerHTML;
      }
      for(const badge of badges.earnedBadges) {
        html.push(/*html*/`
          <a href="https://app.pluralsight.com/profile/sandro-lain" target="_blank" class="badge">
            <!--<div class="badge-img"><img src="${escape(badge.iconUrl)}"/></div>-->
            <div class="badge-title">${escape(badge.name)}</div>
          </a>
        `)
      }
      const skillsHTML = [];
      {
        const response = await fetch("/pluralsight-skills.json");
        const skills = await response.json();
        for(const skill of skills) {
          if(skill.level != "expert") {
            continue;
          }
          skillsHTML.push(/*html*/`
            <a href="https://app.pluralsight.com/profile/sandro-lain" target="_blank" class="skill">
              <!--<div class="skill-img"><img src="${escape(skill.thumbnailUrl)}" /></div>-->
              <div class="skill-name">${escape(skill.title)}</div>
              <div class="skill-lvl">${escape(skill.level)}</div>
              <div class="skill-perc">${escape(skill.percentile)}° percentile</div>
            </a>
          `)
        }
      }
      pluralsightTitle.insertAdjacentHTML("afterend", /*html*/`
        <h3>Pluralsight Skills</h3>
        <div id="skills">${skillsHTML.join("")}</div>
        <h3>Pluralsight Badges</h3>
        <div id="badges">${html.join("")}</div>
      `);
    })();
  }
});



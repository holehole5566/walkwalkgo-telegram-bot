url="http://54.248.150.241:80"

async function renderQuest(id) {
    const root = document.getElementById("root");
    strHTML=""
    quests= await FetchQuest(id);
    strHTML += renderQuestCard(quests)
    root.innerHTML = strHTML
}

async function FetchQuest(id) {
  const response = await fetch(url+"/quest/"+id);
  const jsonData = await response.json();
  return jsonData;
}

function renderQuestOld(id) {
  const root = document.getElementById("root");
  root.innerHTML = `
    <div class="jumbotron">
      <h1 class="display-4">Quest</h1>
      <h1 class="display-4">`+id.toString()+`</h1>
    </div>
  `;
}

function renderFinished(finished) {
  if(finished){
    return "Finished"
  }
  return "Not Finished"
}

function renderQuestCard(quests){
  spots=quests['spots']
  strHTML=`
  <div class="container">
  <div class="row">
    <div class="col-md-6">
      <div class="card">
        <div class="card-body">
          <h3 class="card-title">`+"Year "+quests['year']+" Week "+quests['week']+`</h3>
          <p style="font-size:24px; class="card-text">`+quests['desc']+`</p>
        </div>
      </div>
    </div>
    <div class="col-md-6">
      <div class="card">
        <div class="card-body">
          <h5 class="card-title">`+"Quest 1"+`</h5>
          <p class="card-text">`+spots[0]['name']+`</p>
          <h5 class="card-title">`+"Geo"+`</h5>
          <p class="card-text">`+spots[0]['latitude']+`</p>
          <p class="card-text">`+spots[0]['longitude']+`</p>
          <h5 class="card-title">`+"Status"+`</h5>
          <p class="card-text">`+renderFinished(spots[0]['finished'])+`</p>
        </div>
      </div>
    </div>
    <div class="col-md-6">
      <div class="card">
        <div class="card-body">
          <h5 class="card-title">`+"Quest 2"+`</h5>
          <p class="card-text">`+spots[1]['name']+`</p>
          <h5 class="card-title">`+"Geo"+`</h5>
          <p class="card-text">`+spots[1]['latitude']+`</p>
          <p class="card-text">`+spots[1]['longitude']+`</p>
          <h5 class="card-title">`+"Status"+`</h5>
          <p class="card-text">`+renderFinished(spots[1]['finished'])+`</p>
        </div>
      </div>
    </div>
    <div class="col-md-6">
      <div class="card">
        <div class="card-body">
          <h5 class="card-title">`+"Quest 2"+`</h5>
          <p class="card-text">`+spots[2]['name']+`</p>
          <h5 class="card-title">`+"Geo"+`</h5>
          <p class="card-text">`+spots[2]['latitude']+`</p>
          <p class="card-text">`+spots[2]['longitude']+`</p>
          <h5 class="card-title">`+"Status"+`</h5>
          <p class="card-text">`+renderFinished(spots[2]['finished'])+`</p>
        </div>
      </div>
    </div>
  </div>
</div>
  `
  return strHTML
}

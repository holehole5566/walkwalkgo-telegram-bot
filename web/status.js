url="http://54.248.150.241:80"

async function renderStatus(id,username) {
    const root = document.getElementById("root");
    strHTML=""
    const status = await FetchStatus(id);
    strHTML+=await renderStatusCard(id,username,status)
    root.innerHTML = strHTML;
  }

  async function FetchStatus(id) {
    const response = await fetch(url+"/status/"+id);
    const jsonData = await response.json();
    return jsonData;
  }

  async function renderStatusCard(id,username,status){
    return `
    <div class="container">
    <div class="row justify-content-center">
      <div class="col-md-12">
        <div class="card">
          <div class="card-body">
            <h2 class="card-title">`+"Profile"+`</h2>
            <div class="row">
              <div class="col-md-4">
                <h4 class="card-text">User Name：</h4>
              </div>
              <div class="col-md-8">
                <p class="card-text">`+username+`</p>
              </div>
            </div>
            <div class="row">
              <div class="col-md-4">
                <h4 class="card-text">User ID：</h4>
              </div>
              <div class="col-md-8">
                <p class="card-text">`+id+`</p>
              </div>
            </div>
            <div class="row">
              <div class="col-md-4">
                <h4 class="card-text">Number of Finished Quest：</h4>
              </div>
              <div class="col-md-8">
                <p class="card-text">`+status['count']+`</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
    `
  }




  
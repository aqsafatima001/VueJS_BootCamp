<template>
  <main>
    <!-- <div class="over-lay">
      <div class="modal">
        <textarea name="note" id="note" cols="30" rows="10"></textarea>
        <button>Add Note</button>
        <button id="CloseButton">Close</button>
      </div>
    </div> -->
    <div class="container"> 
      
      <header>
        <h1>Notes</h1>
        <button>+</button>
      </header>

      <div class="cards-container">
        <div class="card">
          <p class="main-text">Lorem ipsum dolor sit amet consectetur adipisicing elit. Ullam consequatur voluptatem sunt ex accusamus, laudantium consequuntur blanditiis tempora.</p>
          <p class="date">04/11/2023</p>
        </div>
        <div class="card">
          <p class="main-text">Lorem ipsum dolor sit amet consectetur adipisicing elit. Ullam consequatur voluptatem sunt ex accusamus, laudantium consequuntur blanditiis tempora.</p>
          <p class="date">04/11/2023</p>
        </div>
      </div>
    </div>
  </main>
</template>

<style scoped>
main{
  height:100vh;
  width: 100vw;

}

.over-lay{
  position: absolute;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0,0, 0.7);
  z-index: 10;
  display: flex;
  align-items: center;
  justify-content: center;

}
.container{
  max-width: 1000px;
  padding: 10px;
  margin: 0 auto;

}
header{
  display: flex;
  justify-content: space-between;
  align-items: center;
}
h1{
  font-weight: bold;
  margin-bottom: 25px;
  font-size: 75px;
}
header button{
  border: none;
  padding: 10px;
  width: 50px;
  height: 50px;
  background-color: rgb(21, 25, 20);
  border-radius: 100%;
  color: white;
  font-size: 20px;
}
.cards-container{
  display: flex;
  flex-wrap: wrap;
}
.card{
  width: 225px;
  height: 225px;
  background-color: rgb(238, 161, 67);
  padding: 10px;
  border-radius: 15px;
  display: flex;
  flex-direction:column ;
  justify-content: space-between;
  margin-right: 20px;
  margin-left: 20px;
}
.date{
  font-size: 12.5px;
  font-weight: bold;
}
.modal{
  width: 750px;
  background-color: white;
  border-radius: 10px;
  padding: 30px;
  position: relative;
  display: flex;
  flex-direction: column;
}
.modal button{
  padding: 10px 20px;
  font-size: 20px;
  width: 100%;
  background-color: blueviolet;
  border: none;
  color: white;
  cursor: pointer;
  margin-top: 15px;

}
#CloseButton{
  background-color: rgb(180, 0, 0);
  margin-top: 7px;
}
</style>
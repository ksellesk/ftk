      url = 'ws://10.14.37.137:8080/ws';
      c = new WebSocket(url);
      
      send = function(data){
        $("#output").append((new Date())+ " ==> "+data+"\n")
        c.send(data)
      }

      c.onmessage = function(msg){
        $("#output").append((new Date())+ " <== "+msg.data+"\n")
        console.log(msg)
      }

      c.onopen = function(){
        setInterval( 
          function(){ send("ping") }
        , 1000 )
      }



export default {
  'GET /serial/list': {
    Ports:[
        "COM3",
        "Broken Port"
    ]
  },
  'GET /serial': {
    Name:null,
    BaudRate:0
  },
  'GET /serial/open': function (req, res) {
    let query = req.query || {};
    if (query.port=="Broken Port"){
      return res.json({
        status:11
      });
    }else if(query.port==null){
      return res.json({
        status:1
      });
    }else {
      return res.json({
        status:0
      });
    }
  },
  'GET /serial/close': {
    status:0
  },

  'GET /variable/types': {
    Types:[
      "double","float","int","int16_t","int32_t","int64_t","int8_t","uint16_t","uint32_t","uint64_t","uint8_t"
    ]
  },
  'GET /file/variables': {
    Variables: [
      {
        Addr:"0x20002020",
        Size:"4",
        Name:"yaw",
        Type:"float"
      },
      {
        Addr:"0x200060be",
        Size:"8",
        Name:"yaw_send",
        Type:"double"
      },
      {
        Addr:"0x2000ac50",
        Size:"4",
        Name:"temp",
        Type:"int"
      },
    ]
  },
  'GET /variable-read/list': {
    Variables: [
      {
        Board:1,
        Name:"yaw",
        Type:"float",
        Addr:"0x20002020",
        Data:0,
        Tick:0
      },
      {
        Board:1,
        Name:"yaw_send",
        Type:"double",
        Addr:"0x200060be",
        Data:0,
        Tick:0
      }
    ]
  },
  'POST /variable-read/add': (req, res) => {
    const { Board, Name, Type, Addr } = req.body;
    if(Name=='bad'){
      return res.json({
        status: 22
      })
    }
    return res.json({
      status: 0
    })
  },
  'POST /variable-read/del': (req, res) => {
    const { Board, Name, Type, Addr } = req.body;
    if(Name=='yaw'){
      return res.json({
        status: 22
      })
    }
    return res.json({
      status: 0
    })
  },
  'GET /variable-modi/list': {
    Variables: [
      {
        Board:1,
        Name:"kp",
        Type:"float",
        Addr:"0x20002070",
        Data:0,
        Tick:0
      },
      {
        Board:1,
        Name:"ki",
        Type:"double",
        Addr:"0x200020a0",
        Data:0,
        Tick:0
      }
    ]
  },
  'POST /variable-modi/add': (req, res) => {
    const { Board, Name, Type, Addr } = req.body;
    return res.json({
      status: 0
    })
  },
  'POST /variable-modi/del': (req, res) => {
    const { Board, Name, Type, Addr } = req.body;
    return res.json({
      status: 0
    })
  },

  'POST /file/upload': (req, res) => {
    return res.json({
      status: 0
    })
  },
  'GET /file/config': {
    IsSaveDataAddr: true,
    IsSaveVariableModi: false,
    IsSaveVariableRead: true
  }
}

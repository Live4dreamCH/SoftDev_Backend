package obj

import (
	"log"

	"github.com/Live4dreamCH/SoftDev_Backend/db"
)

type act struct {
	Aid   int    //`json:"Actid"`
	Uid   int    //"userid"
	Sid  string `json:"SessionID" binding:"required"` 
	Name string `json:"ActName" binding:"required"` 
	Len  int `json:"Length" binding:"required"`
	Des  string `json:"Description" binding:"required"`
	Stop bool //"0:not stop,1:stop"
	Final string  //the final time
	Period []string `json:"OrgPeriods" binding:"required"`
}

func (a *act) Create(aid int,uid int,sid string,name string,len int,des string,stop bool,final string) (suss bool, aid int) {
	
}


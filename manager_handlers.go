package dlg

import (
	"github.com/gin-gonic/gin"
	"github.com/hodgesds/dlg/config"
)

// managerRouter is a Manager HTTP Router.
type managerRouter struct {
	m Manager
}

// NewManagerRouter returns a new manager router.
func NewManagerRouter(e *gin.Engine, m Manager) {
	r := &managerRouter{m: m}
	e.GET("/plans", r.Plans)
	e.GET("/plan/:name", r.Get)
	e.POST("/plan", r.Add)
	e.DELETE("plan/:name", r.Delete)
}

// Plans returns a set of plans.
func (r *managerRouter) Plans(c *gin.Context) {
	plans, err := r.m.Plans(c)
	if err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(200, plans)
}

// Get returns a plan by name.
func (r *managerRouter) Get(c *gin.Context) {
	plan, err := r.m.Get(c, c.Param("name"))
	if err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(200, plan)
}

// Add adds a plan.
func (r *managerRouter) Add(c *gin.Context) {
	var p config.Plan
	if err := c.BindYAML(&p); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	if err := r.m.Add(c, &p); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "ok"})
}

// Delete removes a plan.
func (r *managerRouter) Delete(c *gin.Context) {
	err := r.m.Delete(c, c.Param("name"))
	if err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "ok"})
}

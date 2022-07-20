package ui

import "github.com/tinne26/bindless/src/lang"

type MenuOption struct {
	Text *lang.Text
	SubNode *SubMenu
}

type SubMenu struct {
	selected int
	options []MenuOption
	handler func(string) // option click handler
}

func newSubMenu(options []*lang.Text, handler func(string)) *SubMenu {
	subMenu := &SubMenu {
		selected: -1,
		options: make([]MenuOption, len(options)),
		handler: handler,
	}
	for i, option := range options {
		subMenu.options[i].Text = option
	}
	return subMenu
}

func (self *SubMenu) NavDepth(depth int) *SubMenu {
	if depth  < 0 { return nil  }
	if depth == 0 { return self }
	if self.selected == -1 { return nil }
	subMenu := self.options[self.selected].SubNode
	return subMenu.NavDepth(depth - 1)
}

func (self *SubMenu) NavTo(option string) bool {
	for index, iOption := range self.options {
		if iOption.Text.English() == option {
			self.selected = index
			return true
		}
	}
	return false
}

func (self *SubMenu) Unselect() {
	self.selected = -1
}

func (self *SubMenu) CallHandler() {
	self.handler(self.options[self.selected].Text.English())
}

func (self *SubMenu) SetSubOptions(option string, subOptions []*lang.Text, handler func(string)) bool {
	if len(subOptions) == 0 { panic("can't set empty sub-options") }
	for index, iOption := range self.options {
		if iOption.Text.English() == option {
			self.options[index].SubNode = newSubMenu(subOptions, handler)
			return true
		}
	}
	return false
}

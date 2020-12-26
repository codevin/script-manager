package types

type Item struct {
  Id string
  Name string
  Value string
  Label string
  Option string
  Checked bool
  Disabled bool
  Href string
  Badge string
}

type WidgetInfo struct {
  ScriptName string
  WidgetName string
  WidgetType string
  Template string
  Id string
  Text string
  Body string
  Header string
  HighlightText string
  Items  []Item
  Links []Item
  ButtonText string
  Heading string
  Placeholder string
  Value string
  SelectedValue string
  DefaultValue string
  LogoUrl string
  Title string
  IdEmail string
  IdPassword string
  SubText string
  Label string
  IdDescription string
  Message string
  YesButtonText string
  NoButtonText string
  OkButtonText string
  CancelButtonText string
  PercentCompleted int
  MenuHeight string
  Height string
  Width string
  RememberMe bool
  ShowInline bool 
  ShowSearch bool
  DefaultIsNo bool
  NoCancelButton bool
  CloseWidgetAtEnd bool
  SendOneOutputAtATime bool
}

func CreateWidgetInfo(script_name string) *WidgetInfo {
   wi := WidgetInfo{
      ScriptName: script_name,
      WidgetName: "None",
      WidgetType: "Widget",
      Template: "<div>No Template</div>",
      YesButtonText: "Yes", 
      NoButtonText: "No", 
      OkButtonText: "OK", 
      CancelButtonText: "Cancel", 
   }
   return &wi
}

var WidgetsMap map[string](func(opt *WidgetInfo)(string))

func (wi *WidgetInfo)GetWidgetTemplate() string {
   fn, ok := WidgetsMap[wi.WidgetName]
   if ( ! ok ) {
     return "<div>No template yet for widget: " + wi.WidgetName + "</div>"
   }
   s := fn(wi)
   return s
}

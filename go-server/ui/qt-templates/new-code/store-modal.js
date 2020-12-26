const ModalModule = {
    namespaced: true,
    state: {
      script_id: "",
      widget_name: "",
      key: 0,
    },
    getters: {
    },
    actions: {
       refreshView: (state) => ({
          state.key++
       }),
       poll: (state) => ({
       }),
       AssignWidget(),
       ClearWidget(),
       SendResponse(),
    },
    mutations: {
    },
}

     modal: (modal_id)=>({
         showModal: (modal_id, view) => {
             // Reflect current UI of modal with current state.
             // Pass the Display Vue as 'view' here.
             var modal=this.modals[modal_id]
             view.modalinfo=modal
             view.key++
         },
         showWidget: (widget) => {
             
         },
         hideWidget: () => {
         },
         enablePoll: (script_id) => {
         },
         disablePoll: (script_id) => {
            // Disable previously enabled polling 
         }
     }),

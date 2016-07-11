$(document).ready(function() {

	var writeArea = $(".write-area")[0];
	writeArea.placeholder = '## Clear Markdown';
	writeArea.value = localStorage.getItem('value');

	var getItemContent = localStorage.getItem('display-area');

	if(getItemContent) {
		markedValue(getItemContent);
	}
	
	writeArea.oninput = function() {
		var thisValue = this.value;

		markedValue(thisValue);
		localStorage.setItem('value', thisValue);
		localStorage.setItem('display-area', $('.display-area')[0].innerHTML);
	}

	function markedValue(value) {
		$('.display-area')[0].innerHTML = marked(value);
	}

	marked.setOptions({
		renderer: new marked.Renderer(),
		gfm: true,
		tables: true,
		breaks: true,
		pedantic: false,
		sanitize: true,
		smartLists: true,
		smartypants: false
	});

	// Toolbar 

	$('.toolbar i').on('click', function(event) {
		var targetData = $(event.target).data('tool');
 		event.preventDefault();
 		console.log(targetData);

		switch(targetData) {
			case 'hide-topbar':
				hideTopbar();
				break;
			case 'write-mode':
				writeMode();
				break;
			case 'content-cate':
				showContentCate();
				break;
			case 'change-views':
				changeViews();
				break;
			case 'read-mode':
				readMode();
				break;
		}
	});

	function hideTopbar() {
		console.log('Topbar hidden');
		$('.header').toggleClass('hidden');
		$(event.target).toggleClass('fa-chevron-circle-up')
		.toggleClass('fa-chevron-circle-down');
	}

	function writeMode() {
		console.log('Write Mode');
		$('.markdown-panel').toggleClass('write-mode');
		$(event.target).toggleClass('fa-desktop')
		.toggleClass('fa-columns');
	}

	function readMode() {
		console.log('Read Mode');
		$('.markdown-panel').toggleClass('read-mode');
		$(event.target).toggleClass('fa-desktop')
		.toggleClass('fa-columns');
	}

	function showContentCate() {
		console.log('Content Cate Showed');
	}

	function changeViews() {
		console.log('View Changed');
		$('.markdown-panel').toggleClass('change-views');
		$(event.target).toggleClass('fa-chevron-circle-left')
		.toggleClass('fa-chevron-circle-right');
	}
});
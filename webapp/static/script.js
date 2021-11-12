function ani() {
	  var sr = document.getElementById('search-gists')
	  var sbyu = []

      fetch('https://sulphurous.cf/api/gists/all/')
		  .then(r=>r.json()).then(r=>
		  {
			r['gists'].forEach(g => {
				if (g['user'] === sr.value || g['title'] === sr.value || g['id'].toString() === sr.value)
					sbyu.push(g)
			})
			if (sbyu.length === 0) 
      		{
				document.getElementById('status').innerText = 'No results. None at all.'
        		document.getElementsByClassName('loading')[0].hidden = true
      		} else 
      		{
        		sbyu.forEach(sg => {
					fetch(`https://scratchapinc.herokuapp.com/users/${sg['user']}/`)
					.then(r=>r.json()).then(r=>
					{
          				document.getElementById('results').innerHTML += `<img src='${r['profile']['images']['55x55']}' height='25' width='25'> <a href='https://sulphurous.cf/gists/${sg['id']}' target='blank'>${sg['title']} by ${sg['user']}</a> <br>`
					})
        		})
            document.getElementById('status').innerText = 'Found something.'
        		document.getElementsByClassName('loading')[0].hidden = true
            document.getElementById('results').hidden = false
      		}
	})
}

function toggleAccordion(accordionId) {
	const accordion = document.getElementById(accordionId);
	const icon = document.getElementById(`icon${accordionId}`);

	if (accordion.style.display === 'block') {
	  accordion.style.display = 'none';
	  icon.classList.remove('rotate-180');
	} else {
	  // Close other opened accordions
	  const accordions = document.getElementsByClassName('accordion');
	  const icons = document.getElementsByClassName('accordion-icon');

	  for (let i = 0; i < accordions.length; i++) {
		const acc = accordions[i];
		const icn = icons[i];

		if (acc.id !== accordionId && acc.style.display === 'block') {
		  acc.style.display = 'none';
		  icn.classList.remove('rotate-180');
		}
	  }
	  accordion.style.display = 'block';
	  icon.classList.add('rotate-180');
	}
  }

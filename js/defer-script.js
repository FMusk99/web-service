function handleLoad(e) {
    console.log('Loaded import: ' + e.target.href);
  }

function handleError(e) {
    console.log('Error loading import: ' + e.target.href);
}
const fetch = require('node-fetch');

async function updateBranchStatus(branch, commit, token) {
  const [owner, repo] = process.env.TRAVIS_REPO_SLUG.split('/');
  const url = `http://${owner}.github.io/${repo}/preview/${branch}/`;
  const res = await fetch(`https://api.github.com/repos/${owner}/${repo}/statuses/${commit}`, {
    method: 'POST',
    headers: {
      Authorization: `token ${token}`,
      'User-Agent': 'swagger-repo-travis'
    },
    body: JSON.stringify({
      state: 'success',
      target_url: url,
      description: 'Link to preview',
      context: 'Preview'
    })
  });

  if (!res.ok || res.status !== 201) {
    throw new Error(await res.text());
  }
}

exports.notifyBranchPreviewFromTravis = async function(branch, commit) {
  try {
    await updateBranchStatus(branch, commit, process.env.GH_TOKEN);
  } catch (e) {
    console.log('Failed to update branch status on GitHub:' + e.message);
  }
};

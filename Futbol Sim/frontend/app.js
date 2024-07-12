document.getElementById("simulateButton").addEventListener("click", simulate);
document.getElementById("nextWeekButton").addEventListener("click", nextWeek);
window.addEventListener("load", loadStandings);
window.addEventListener("load", loadMatches);

function simulate() {
  fetch("/simulate")
    .then((response) => response.json())
    .then((data) => {
      loadStandings();
      loadMatches();
    })
    .catch((error) => console.error("Error:", error));
}

function nextWeek() {
  fetch("/next-week")
    .then((response) => response.json())
    .then((data) => {
      loadStandings();
      loadMatches();
    })
    .catch((error) => console.error("Error:", error));
}

function loadStandings() {
  fetch("/standings")
    .then((response) => response.json())
    .then((data) => {
      const tbody = document.querySelector("#standingsTable tbody");
      tbody.innerHTML = "";
      data.forEach((team) => {
        const row = document.createElement("tr");
        row.innerHTML = `
          <td>${team.Name}</td>
          <td>${team.Points}</td>
          <td>${team.Wins}</td>
          <td>${team.Draws}</td>
          <td>${team.Losses}</td>
          <td>${team.GoalsFor}</td>
          <td>${team.GoalsAgainst}</td>
        `;
        tbody.appendChild(row);
      });
    })
    .catch((error) => console.error("Error:", error));
}

function loadMatches() {
  fetch("/matches")
    .then((response) => response.json())
    .then((data) => {
      const tbody = document.querySelector("#matchesTable tbody");
      tbody.innerHTML = "";
      if (data && Array.isArray(data)) {
        // Null veya undefined kontrolü
        data.forEach((match) => {
          const row = document.createElement("tr");
          row.innerHTML = `
            <td>${match.home_team}</td>
            <td>${match.away_team}</td>
            <td>${match.home_goals}</td>
            <td>${match.away_goals}</td>
          `;
          tbody.appendChild(row);
        });
      } else {
        console.error("Invalid data format for matches");
      }
    })
    .catch((error) => console.error("Error:", error));
}

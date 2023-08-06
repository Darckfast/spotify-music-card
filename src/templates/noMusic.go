package templates

var NoMusic = `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="480" height="200">
    <foreignObject width="480" height="200">
         <style>
            .song {
                color: #9b9b9b;
                margin: 0.7em 0 0.7em 0;
                font-size: 24px;
                width: 13rem;
                mix-blend-mode: difference;

                text-align: center;
                white-space: nowrap;
                text-overflow: ellipsis;
            }

            .artist {
                color: #c9c9c9;
                font-size: 20px;
                margin: 5px 0 0 0;
                mix-blend-mode: difference;

                text-align: center;
            }

            .bars {
                height: 3px;

                display: flex;
                justify-content: center;
                align-items: center;

                .bar {
                    width: 4px;

                    background: #8257e6;
                    margin: 0.5px;
                    animation: sound 0ms -400ms linear infinite alternate;
                }

                .bar-no-animation {
                    width: 4px;

                    background: #8257e6;
                    margin: 0.5px;
                }
            }

            .music-eq {
                display: flex;
                justify-content: center;
                align-items: center;
                flex-direction: column; 
                position: relative;
            }

            .music-eq::after {
                content: '';
                height: 100%;
                width: 100%;
                position: absolute;
                /* From https://css.glass */
                background: rgba(255, 255, 255, 0.25);
                border-radius: 8px;
                box-shadow: 0 4px 30px rgba(0, 0, 0, 0.1);
                backdrop-filter: blur(10px);
                -webkit-backdrop-filter: blur(10px);
                border: 1px solid rgba(255, 255, 255, 0.33);

            }

            .main-card {
                display: flex;
                justify-content: space-evenly;
                align-items: center;
                font-family: Roboto, sans-serif;
            }
        </style>
          <div xmlns="http://www.w3.org/1999/xhtml">
            <div class="main-card"><a><img src="https://i.imgur.com/FAMtjqN.png" /></a>
                <div class="music-eq">
                    <div class="song">No music playing now</div>
                    <div class="artist">None</div>
                    <div class="bars">
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                        <div class="bar-no-animation"></div>
                    </div>
                </div>
            </div>
        </div>
    </foreignObject>
</svg>`

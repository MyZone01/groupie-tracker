class DateRangeSlider extends HTMLElement {
    constructor() {
        super();
        this.attachShadow({ mode: 'open' });
        this.minGap = 0;
        this.sliderMinValue = 1950;
        this.sliderMaxValue = 2023;
        this.sliderName = this.getAttribute('name') || 'date-slider';
    }

    connectedCallback() {
        this.render();
        this.minSlider = this.shadowRoot.getElementById(`min-${this.sliderName}`);
        this.maxSlider = this.shadowRoot.getElementById(`max-${this.sliderName}`);
        this.displayValOne = this.shadowRoot.getElementById(`min-${this.sliderName}-value`);
        this.displayValTwo = this.shadowRoot.getElementById(`max-${this.sliderName}-value`);
        this.sliderTrack = this.shadowRoot.getElementById(`${this.sliderName}-track`);
        this.minSlider.addEventListener('input', () => this.slide());
        this.maxSlider.addEventListener('input', () => this.slide());
        this.minSlider.addEventListener('change', this.emitChangeEvent.bind(this));
        this.maxSlider.addEventListener('change', this.emitChangeEvent.bind(this));
        this.slide();
        this.slide();
    }

    slide() {
        if (parseInt(this.maxSlider.value) - parseInt(this.minSlider.value) <= this.minGap) {
            this.minSlider.value = parseInt(this.maxSlider.value) - this.minGap;
        }
        this.displayValOne.textContent = this.minSlider.value;

        if (parseInt(this.maxSlider.value) - parseInt(this.minSlider.value) <= this.minGap) {
            this.maxSlider.value = parseInt(this.minSlider.value) + this.minGap;
        }
        this.displayValTwo.textContent = this.maxSlider.value;
        this.fillColor();
    }

    emitChangeEvent() {
        let changeEvent = new CustomEvent('change', {
            detail: {
                minValue: this.minValue,
                maxValue: this.maxValue
            }
        });
        this.dispatchEvent(changeEvent);
    }

    get name() {
        return this.getAttribute('name');
    }

    get minValue() {
        return this.minSlider.value;
    }

    get maxValue() {
        return this.maxSlider.value;
    }

    fillColor() {
        let sliderOneValue = this.minSlider.value ?? this.sliderMinValue;
        let percent1 = ((sliderOneValue - this.sliderMinValue) / (this.sliderMaxValue - this.sliderMinValue)) * 100;

        let sliderTwoValue = this.maxSlider.value ?? this.sliderMaxValue;
        let percent2 = ((sliderTwoValue - this.sliderMinValue) / (this.sliderMaxValue - this.sliderMinValue)) * 100;

        let sliderProgress = `linear-gradient(to right, #dadae5 ${percent1}% , #3264fe ${percent1}% , #3264fe ${percent2}%, #dadae5 ${percent2}%)`;
        this.sliderTrack.style.background = sliderProgress;
    }

    render() {
        this.shadowRoot.innerHTML = /*html*/`
        <style>
        .root {
            display: flex;
            align-items: center;
            justify-content: center;
        }

        .wrapper {
            position: relative;
            width: 150px;
        }

        .slider {
            position: relative;
            width: 100%;
            height: 25px;
        }

        input[type="range"] {
            -webkit-appearance: none;
            -moz-appearance: none;
            appearance: none;
            width: 100%;
            outline: none;
            position: absolute;
            margin: auto;
            top: 0;
            bottom: 0;
            background-color: transparent;
            pointer-events: none;
        }

        [id$='-track'] {
            width: 100%;
            height: 2px;
            position: absolute;
            margin: auto;
            top: 0;
            bottom: 0;
            border-radius: 5px;
        }

        input[type="range"]::-webkit-slider-runnable-track {
            -webkit-appearance: none;
            height: 2px;
        }

        input[type="range"]::-moz-range-track {
            -moz-appearance: none;
            height: 2px;
        }

        input[type="range"]::-ms-track {
            appearance: none;
            height: 2px;
        }

        input[type="range"]::-webkit-slider-thumb {
            -webkit-appearance: none;
            height: .75em;
            width: .75em;
            background-color: #3264fe;
            cursor: pointer;
            margin-top: -4px;
            pointer-events: auto;
            border-radius: 50%;
        }

        input[type="range"]::-moz-range-thumb {
            -webkit-appearance: none;
            height: .75em;
            width: .75em;
            cursor: pointer;
            border-radius: 50%;
            background-color: #3264fe;
            pointer-events: auto;
        }

        input[type="range"]::-ms-thumb {
            appearance: none;
            height: .75em;
            width: .75em;
            cursor: pointer;
            border-radius: 50%;
            background-color: #3264fe;
            pointer-events: auto;
        }

        input[type="range"]:active::-webkit-slider-thumb {
            background-color: #ffffff;
            border: 3px solid #3264fe;
        }

        .values {
            background-color: #3264fe;
            position: relative;
            margin: auto;
            padding: 1px 0;
            border-radius: 5px;
            text-align: center;
            font-weight: 500;
            color: #ffffff;
        }

        .values:before {
            content: "";
            position: absolute;
            height: 0;
            width: 0;
            border-top: 7.5px solid #3264fe;
            border-left: 7.5px solid transparent;
            border-right: 7.5px solid transparent;
            margin: auto;
            bottom: -7.5px;
            left: 0;
            right: 0;
        }
    </style>
        <div class="root">
            ${this.sliderName}
            <div class="wrapper">
                <div class="values">
                    <span id="min-${this.sliderName}-value">
                        ${this.sliderMinValue}
                    </span>
                    <span> &dash; </span>
                    <span id="max-${this.sliderName}-value">
                        ${this.sliderMaxValue}
                    </span>
                </div>
                <div class="slider">
                    <div id="${this.sliderName}-track"></div>
                    <input type="range" step="1" name="min-${this.sliderName}" min="${this.sliderMinValue}" max="${this.sliderMaxValue}" value="${this.getAttribute('min-value') || this.sliderMinValue}" id="min-${this.sliderName}" />
                    <input type="range" step="1" name="max-${this.sliderName}" min="${this.sliderMinValue}" max="${this.sliderMaxValue}" value="${this.getAttribute('max-value') || this.sliderMaxValue}" id="max-${this.sliderName}" />
                </div>
            </div>
        </div>
        `;
    }
}

customElements.define('date-range-slider', DateRangeSlider);
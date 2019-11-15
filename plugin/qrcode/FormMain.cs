using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace qrcode
{
    public partial class FormMain : Form
    {
        public FormMain()
        {
            InitializeComponent();
        }

        private void FormMain_Load(object sender, EventArgs e)
        {
            // FormMain Initial
            AutoScaleMode = AutoScaleMode.Dpi;
            MaximizeBox = false;
            MinimizeBox = true;
            DoubleBuffered = true;
            StartPosition = FormStartPosition.CenterScreen;

            // QRCode
            this.textBox_content.ReadOnly = true;
        }

        private void FormMain_FormClosed(object sender, FormClosedEventArgs e)
        {

        }

        private void Button_generate_Click(object sender, EventArgs e)
        {

        }

        private void Button_clear_Click(object sender, EventArgs e)
        {

        }

        private void Button_select_Click(object sender, EventArgs e)
        {

        }
    }
}
